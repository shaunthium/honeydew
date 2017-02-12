package deploy

func Deploy(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname string) error {
	// Build image in current directory
	if err := BuildImage(imageTagName); err != nil {
		return err
	}
	if err := PushImage(imageTagName); err != nil {
		return err
	}
	if err := RunCommandInServer(volumeMountHostname, volumeName, targetDirectory, imageTagName, instanceHostname); err != nil {
		return err
	}
	return nil
}
