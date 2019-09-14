package mattermost

import (
	"fmt"
	"os"

	"github.com/mattermost/mattermost-server/model"
)

// NewMatterMostClient returns a NewAPIV4Client after logging in with the provided username and password
func NewMatterMostClient(url string, username string, password string) *model.Client4 {
	api := model.NewAPIv4Client(url)

	api.Login(username, password)

	return api
}

// GetChannelMembers retrieves a list of members in a given channel for the specified teamName
func GetChannelMembers(m model.Client4, teamName string, channelName string) *model.ChannelMembers {
	team, resp := m.GetTeamByName(teamName, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", team)

	channel, resp := m.GetChannelByName(channelName, team.Id, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	//fmt.Printf("%+v\n", channel)

	members, resp := m.GetChannelMembers(channel.Id, 0, 100, "")
	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error: %+v", resp)
		os.Exit(1)
	}
	return members
}
