import sys
import datetime
import hashlib
import os
import json
import random
from typing import Optional, Dict, Any

from .hash import random_hash
from .metadata import rfc3339_datetime
from .storage import Storage


# this slows down import replicate considerably, but is probably fine
# since the framework is probably loaded/about to be loaded anyway
# FIXME (bfirsh): if the user has both tensorflow and numpy installed, then
# they'll unnecessarily loading both.
# fmt: off
try:
    import numpy as np  # type: ignore
    has_numpy = True
except ImportError:
    has_numpy = False
try:
    import torch  # type: ignore
    has_torch = True
except ImportError:
    has_torch = False
try:
    import tensorflow as tf  # type: ignore
    has_tensorflow = True
except ImportError:
    has_tensorflow = False
# fmt: on


class CustomJSONEncoder(json.JSONEncoder):
    def default(self, obj):
        if has_numpy:
            if isinstance(obj, np.integer):
                return int(obj)
            elif isinstance(obj, np.floating):
                return float(obj)
            elif isinstance(obj, np.ndarray):
                return obj.tolist()
        if has_torch and isinstance(obj, torch.Tensor):
            return obj.detach().tolist()
        if has_tensorflow and isinstance(obj, tf.Tensor):
            return obj.numpy().tolist()
        print(type(obj))
        return json.JSONEncoder.default(self, obj)


class Commit(object):
    """
    A snapshot of a training job -- the working directory plus any metadata.
    """

    def __init__(
        self,
        experiment,  # can't type annotate due to circular import
        path: Optional[str],
        project_dir: str,
        created: datetime.datetime,
        step: Optional[int],
        labels: Dict[str, Any],
    ):
        self.experiment = experiment
        self.project_dir = project_dir
        self.path = path
        self.created = created
        self.step = step
        self.labels = labels

        # TODO (bfirsh): content addressable id
        self.id = random_hash()

        self.validate_labels()

    def save(self, storage: Storage):
        if self.path is not None:
            source_path = os.path.join(self.project_dir, self.path)
            destination_path = os.path.join("commits", self.id, self.path)
            if os.path.isfile(source_path):
                with open(os.path.join(source_path), "rb") as fh:
                    data = fh.read()
                storage.put(destination_path, data)
            else:
                storage.put_directory(destination_path, source_path)
        obj = {
            "id": self.id,
            "created": rfc3339_datetime(self.created),
            "experiment_id": self.experiment.id,
            "path": self.path,
            "labels": self.labels,
        }
        if self.step is not None:
            obj["step"] = self.step
        storage.put(
            "metadata/commits/{}.json".format(self.id),
            json.dumps(obj, indent=2, cls=CustomJSONEncoder),
        )

    def validate_labels(self):
        metrics = self.experiment.config.get("metrics", [])
        metric_keys = set(
            filter(lambda x: x, [metric.get("name") for metric in metrics])
        )
        label_keys = set(self.labels.keys())
        missing_keys = metric_keys - label_keys
        if missing_keys:
            print(
                "Warning: You specified these metrics in replicate.yaml, but they are missing in your call to replicate.commit(): {}".format(
                    ", ".join(missing_keys)
                ),
                file=sys.stderr,
            )
