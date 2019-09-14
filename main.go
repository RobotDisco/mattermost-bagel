package main

import (
	"fmt"

	"github.com/mattermost/mattermost-server/model"
)

func main() {
	c := model.NewAPIv4Client("https://mattermost.internal.tulip.io")
	fmt.Println(c)
}
