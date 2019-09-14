package main

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-server/model"
)

func main() {
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")

	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	teamName := os.Getenv("BAGEL_TEAM_NAME")
	channelName := os.Getenv("BAGEL_CHANNEL_NAME")

	api := model.NewAPIv4Client(serverURL)
	api.Login(botUserName, botPassword)
	team, err := api.GetTeamByName(teamName, "")
	if err.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", err)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", team)

	channel, err := api.GetChannelByName(channelName, team.Id, "")
	if err.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", err)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", channel)

	members, err := api.GetChannelMembers(channel.Id, 0, 100, "")
	if err.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", members)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(*members), channelName, teamName)
}
