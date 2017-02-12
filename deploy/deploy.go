package deploy

import (
	"fmt"
)

func Deploy(volumeMountHostname, volumeName, targetDirectory string) error {
	fmt.Println("in deploy")
	if err := RunCommandInServer(volumeMountHostname, volumeName, targetDirectory); err != nil {
		return err
	}
	// ListImages()
	// if err := BuildImage(); err != nil {
	// 	return err
	// }
	return nil
}
