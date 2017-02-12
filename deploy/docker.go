package deploy

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/net/context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func BuildImage() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	if err := tarFolder(cwd, "/tmp"); err != nil {
		return err
	}

	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}

	tarFilepath := "/" + filepath.Join("tmp", filepath.Base(cwd)) + ".tar"
	buildCtx, err := os.Open(tarFilepath)
	defer buildCtx.Close()
	if err != nil {
		return err
	}
	buildOptions := types.ImageBuildOptions{
		Tags:       []string{"hi"},
		Dockerfile: "sample-app-1/Dockerfile",
	}
	reader, err := cli.ImageBuild(context.Background(), buildCtx, buildOptions)
	if err != nil {
		return err
	}
	b := []byte{}
	reader.Body.Read(b)
	defer reader.Body.Close()
	fmt.Printf("b: %v\n", b)
	return nil
}

func ListImages() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Printf("%v\n", container)
	}
}
