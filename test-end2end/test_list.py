import json
import os
import subprocess
import pytest  # type: ignore


@pytest.mark.parametrize("storage_backend", ["s3", "file", "undefined"])
def test_list(storage_backend, tmpdir, temp_bucket, tmpdir_factory):
    if storage_backend == "s3":
        storage = "s3://" + temp_bucket
    elif storage_backend == "file":
        storage = "file://" + str(tmpdir_factory.mktemp("storage"))
    elif storage_backend == "undefined":
        storage = str(tmpdir_factory.mktemp("storage"))

    with open(os.path.join(tmpdir, "replicate.yaml"), "w") as f:
        f.write(
            """
storage: {storage}
""".format(
                storage=storage
            )
        )
    with open(os.path.join(tmpdir, "train.py"), "w") as f:
        f.write(
            """
import replicate

def main():
    experiment = replicate.init(params={"my-param": "my-value"})

    for step in range(3):
        experiment.commit(metrics={"step": step})

if __name__ == "__main__":
    main()
"""
        )

    env = os.environ
    env["REPLICATE_NO_ANALYTICS"] = "1"
    env["PATH"] = "/usr/local/bin:" + os.environ["PATH"]

    return_code = subprocess.Popen(
        ["python", "train.py", "train.py"], cwd=tmpdir, env=env,
    ).wait()
    assert return_code == 0

    experiments = json.loads(
        subprocess.check_output(
            ["replicate", "list", "--format=json"], cwd=tmpdir, env=env,
        )
    )
    assert len(experiments) == 1

    exp = experiments[0]
    latest = exp["latest_commit"]
    assert len(exp["id"]) == 64
    assert exp["params"] == {"my-param": "my-value"}
    assert exp["num_commits"] == 3
    assert len(latest["id"]) == 64
    # FIXME: now rfc3339 strings
    # assert latest["created"] > exp["created"]
    assert latest["experiment"]["id"] == exp["id"]
    assert latest["metrics"] == {"step": 2}
