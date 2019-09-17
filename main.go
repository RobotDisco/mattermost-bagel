package main

import (
	"fmt"
	"os"

	"github.com/RobotDisco/mattermost-bagel/mattermost"
)

func main() {
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")

	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	teamName := os.Getenv("BAGEL_TEAM_NAME")
	channelName := os.Getenv("BAGEL_CHANNEL_NAME")

	api := mattermost.NewMatterMostClient(serverURL, botUserName, botPassword)

	members := mattermost.GetActiveChannelMembers(*api, teamName, channelName)
	fmt.Printf("%+v\n", members)
	fmt.Printf("There are %d members in channel %s for team %s\n", len(members), channelName, teamName)
	bot := mattermost.GetBotUser(*api)
	pairs := mattermost.SplitIntoPairs(members, bot.Id)
	fmt.Printf("%+v\n", pairs[0].First)

	mattermost.MessageMembers(*api, pairs, bot)
}
