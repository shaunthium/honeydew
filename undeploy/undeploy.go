package undeploy

import (
	"os"

	"github.com/shaunthium/honeydew/ssh"
)

func Undeploy(targetDirectory, instanceHostname string) error {
	return undeployFromServer(targetDirectory, instanceHostname)
}

func undeployFromServer(targetDirectory, instanceHostname string) error {
	dockerKillCmd := &ssh.SSHCommand{
		Path:   "sudo docker kill $(sudo docker ps -q)",
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	unmountCmd := &ssh.SSHCommand{
		Path:   "sudo umount " + targetDirectory,
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	cmds := []*ssh.SSHCommand{dockerKillCmd, unmountCmd}

	return ssh.RunCommandsInServer(instanceHostname, cmds)
}
