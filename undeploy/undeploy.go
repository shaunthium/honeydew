package undeploy

import (
	"os"

	"github.com/shaunthium/honeydew/ssh"
)

func Undeploy(targetDirectory, instanceHostname, dockerVolumeName string) error {
	return undeployFromServer(targetDirectory, instanceHostname, dockerVolumeName)
}

func undeployFromServer(targetDirectory, instanceHostname, dockerVolumeName string) error {
	unmountCmd := &ssh.SSHCommand{
		Path:   "sudo umount " + targetDirectory,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	dockerKillCmd := &ssh.SSHCommand{
		Path:   "sudo docker kill $(sudo docker ps -aq)",
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	dockerRmCmd := &ssh.SSHCommand{
		Path:   "sudo docker rm $(sudo docker ps -aq)",
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	dockerUnmountVolumeCmd := &ssh.SSHCommand{
		Path:   "sudo docker volume rm " + dockerVolumeName,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmds := []*ssh.SSHCommand{unmountCmd, dockerKillCmd, dockerRmCmd, dockerUnmountVolumeCmd}

	return ssh.RunCommandsInServer(instanceHostname, cmds)
}
