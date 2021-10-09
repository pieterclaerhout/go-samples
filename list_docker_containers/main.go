package main

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = false

	cli, err := client.NewEnvClient()
	log.CheckError(err)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	log.CheckError(err)

	for _, container := range containers {
		log.InfoDump(container, container.Image)
	}

}
