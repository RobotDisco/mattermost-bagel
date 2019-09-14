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

	members := GetChannelMembers(*api, teamName, channelName)
	bot := GetBotUser(*api)

	memberCount := len(*members) - 1

	pairs := SplitIntoPairs(*members, bot.Id)
	pairCount := len(pairs)
	if pairCount != (memberCount / 2) {
		t.Errorf("Expected %d; Actual %d\nMembers: %v\nPairs: %v\n", memberCount/2, pairCount, members, pairs)
	}
}

// TODO: Test bot user is filtered out
// TODO: Are users different (based on history)?
// TODO: Are users disabled?
// TODO: Some kind of memory, have these users been matched recently?
// TODO: What do we do with the odd person out? (Sorry you're not interesting, here's a haiku?)
