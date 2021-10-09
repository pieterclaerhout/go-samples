package main

import (
	"context"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pieterclaerhout/go-formatter"
	"github.com/pieterclaerhout/go-log"
)

func main() {

	log.PrintColors = true
	log.PrintTimestamp = false

	cli, err := client.NewEnvClient()
	log.CheckError(err)

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{
		All: false,
	})

	for _, image := range images {
		for _, tag := range image.RepoTags {

			tagParts := strings.SplitN(tag, ":", 2)

			log.Infof(
				"%-25s | %-15s | %s | %-30s | %10s",
				tagParts[0],
				tagParts[1],
				image.ID[7:19],
				time.Unix(image.Created, 0),
				formatter.FileSize(image.Size),
			)
		}
	}

}
