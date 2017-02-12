package deploy

import (
	"fmt"
	"os/exec"
)

func BuildImage(imageTagName string) error {
	cmd := exec.Command("docker", "build", "-t", imageTagName, ".")
	fmt.Println("Building image: " + imageTagName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func PushImage(imageTagName string) error {
	fmt.Println("Pushing image: " + imageTagName)
	cmd := exec.Command("docker", "push", imageTagName)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
