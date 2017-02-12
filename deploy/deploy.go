package deploy

import (
	"os"

	"github.com/shaunthium/honeydew/ssh"
)

func Deploy(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname, dockerVolumeName string, isApi bool) error {
	if !isApi {
		// Build image in current directory and push
		if err := BuildImage(imageTagName); err != nil {
			return err
		}
		if err := PushImage(imageTagName); err != nil {
			return err
		}
	}
	if err := deployToServer(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname, dockerVolumeName); err != nil {
		return err
	}
	return nil
}

func deployToServer(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname, dockerVolumeName string) error {
	mountCmd := &ssh.SSHCommand{
		Path:   "sudo mount " + volumeMountHostname + ":/" + volumeName + " " + targetDirectory,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	dockerMountVolumeCmd := &ssh.SSHCommand{
		Path:   "sudo docker volume create -d netapp --name " + dockerVolumeName,
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
		Path:   "sudo docker run --rm -it --volume-driver netapp --volume " + dockerVolumeName + ":/sample-app-1 " + imageTagName,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmds := []*ssh.SSHCommand{mountCmd, dockerMountVolumeCmd, dockerPullCmd, dockerRunCmd}

	return ssh.RunCommandsInServer(instanceHostname, cmds)
}
