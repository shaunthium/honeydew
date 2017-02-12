package deploy

import "fmt"

func Deploy(volumeMountHostname, volumeName, targetDirectory string) error {
	fmt.Println("in deploy")
	RunCommandInServer(volumeMountHostname, volumeName, targetDirectory)
	return nil
}
