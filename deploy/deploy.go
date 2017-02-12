package deploy

import (
	"fmt"
)

func Deploy(volumeMountHostname, volumeName, targetDirectory, imageTagName string) error {
	fmt.Println("in deploy")

	// Build image in current directory
	if err := BuildImage(imageTagName); err != nil {
		return err
	}
	if err := PushImage(imageTagName); err != nil {
		return err
	}
	// if err := RunCommandInServer(volumeMountHostname, volumeName, targetDirectory); err != nil {
	// 	return err
	// }
	return nil
}
