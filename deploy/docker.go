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

// func BuildImage() error {
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}
// 	if err := tarFolder(cwd, "/tmp"); err != nil {
// 		return err
// 	}
//
// 	cli, err := client.NewEnvClient()
// 	if err != nil {
// 		return err
// 	}
//
// 	tarFilepath := "/" + filepath.Join("tmp", filepath.Base(cwd)) + ".tar"
// 	buildCtx, err := os.Open(tarFilepath)
// 	defer buildCtx.Close()
// 	if err != nil {
// 		return err
// 	}
// 	buildOptions := types.ImageBuildOptions{
// 		Tags:       []string{"shaunthium/honeydew:v2"},
// 		Dockerfile: "sample-app-1/Dockerfile",
// 	}
// 	reader, err := cli.ImageBuild(context.Background(), buildCtx, buildOptions)
// 	if err != nil {
// 		return err
// 	}
// 	b := []byte{}
// 	reader.Body.Read(b)
// 	defer reader.Body.Close()
// 	fmt.Printf("b: %v\n", b)
// 	return nil
// }
