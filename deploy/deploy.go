package deploy

import (
	"fmt"

	"github.com/shaunthium/honeydew/ssh"
)

func Deploy(volumeMountHostname, volumeName, targetDirectory string) error {
	fmt.Println("in deploy")
	ssh.RunCommandInServer(volumeMountHostname, volumeName, targetDirectory)
	return nil
}
