package mattermost

import (
	"fmt"
	"os"
	"testing"
)

func TestNumberOfPairs(t *testing.T) {
	fmt.Println("TestNumberOfPairs")
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")

	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	teamName := os.Getenv("BAGEL_TEAM_NAME")
	channelName := os.Getenv("BAGEL_CHANNEL_NAME")

	api := NewMatterMostClient(serverURL, botUserName, botPassword)

	channelID, members := GetActiveChannelMembers(*api, teamName, channelName)
	bot := GetBotUser(*api)

	memberCount := len(members) - 1

	pairs := SplitIntoPairs(channelID, members, bot.Id)
	pairCount := len(pairs)
	if pairCount != (memberCount / 2) {
		t.Errorf("Expected %d; Actual %d\nMembers: %v\nPairs: %v\n", memberCount/2, pairCount, members, pairs)
	}
}
