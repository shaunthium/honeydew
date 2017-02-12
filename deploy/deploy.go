package deploy

import (
	"os"

	"github.com/shaunthium/honeydew/ssh"
)

func Deploy(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname string) error {
	// Build image in current directory
	if err := BuildImage(imageTagName); err != nil {
		return err
	}
	if err := PushImage(imageTagName); err != nil {
		return err
	}
	if err := deployToServer(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname); err != nil {
		return err
	}
	return nil
}

func deployToServer(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname string) error {
	mountCmd := &ssh.SSHCommand{
		Path:   "sudo mount " + volumeMountHostname + ":/" + volumeName + " " + targetDirectory,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	dockerPullCmd := &ssh.SSHCommand{
		Path:   "sudo docker pull " + imageTagName,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}

	dockerRunCmd := &ssh.SSHCommand{
		Path:   "sudo docker run " + imageTagName,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmds := []*ssh.SSHCommand{mountCmd, dockerPullCmd, dockerRunCmd}

	return ssh.RunCommandsInServer(instanceHostname, cmds)
}
