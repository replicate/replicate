package docker

import (
	"io"
	"os"
	"os/exec"

	"replicate.ai/cli/pkg/console"
	"replicate.ai/cli/pkg/remote"
)

// Build a Docker image by calling `docker build` locally or remotely over SSH
//
// Log output is sent to stdout/err.
func Build(remoteOptions *remote.Options, folder string, dockerfile string, name string, baseImage string, hasGPU bool) error {
	args := []string{
		"build", ".",
		"--build-arg", "BUILDKIT_INLINE_CACHE=1",
		"--build-arg", "BASE_IMAGE=" + baseImage,
		"--file", "-",
		"--tag", name,
	}
	// TODO(andreas): detect if terminal supports cursor movement
	if !console.IsTTY() {
		args = append(args, "--progress", "plain")
	}
	if hasGPU {
		args = append(args, "--build-arg", "HAS_GPU=1")
	}

	cmd := exec.Command("docker", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, dockerfile) //nolint
	}()

	console.Debug("Running '%s'", cmd.String())

	// Local
	if remoteOptions == nil {
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "DOCKER_BUILDKIT=1")
		cmd.Dir = folder
		if err := cmd.Start(); err != nil {
			return err
		}
		return cmd.Wait()
	}

	// Remote, via SSH
	cmd.Env = os.Environ() // the environment gets filtered later in client.WrapCommandSafeEnv
	cmd.Env = append(cmd.Env, "DOCKER_BUILDKIT=1")

	remoteTempDir, err := remote.UploadToTempDir(folder, remoteOptions)
	if err != nil {
		return err
	}
	cmd.Dir = remoteTempDir
	client, err := remote.NewClient(remoteOptions)
	if err != nil {
		return err
	}
	remoteCmd, err := client.WrapCommandSafeEnv(cmd)
	if err != nil {
		return err
	}
	if err := remoteCmd.Start(); err != nil {
		return err
	}
	return remoteCmd.Wait()
}
