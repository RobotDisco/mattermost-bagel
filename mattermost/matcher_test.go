package mattermost

import (
	"fmt"
	"testing"
	"github.com/RobotDisco/mattermost-bagel/config"
)

func TestNumberOfPairs(t *testing.T) {
	cfg := config.Config{
		MattermostURL: "http://localhost:8065",
		MattermostUser: "coffeebot1",
		MattermostPassword: "password",
		MattermostTeam: "test",
		MattermostChannel: "coffeebot",
	}
	
	fmt.Println("TestNumberOfPairs")
	api := NewMatterMostClient(cfg.MattermostURL, cfg.MattermostUser, cfg.MattermostPassword)

	channelID, members := GetActiveChannelMembers(*api, cfg.MattermostTeam, cfg.MattermostChannel)
	bot := GetBotUser(*api)

	memberCount := len(members) - 1

	pairs := SplitIntoPairs(channelID, members, bot.Id)
	pairCount := len(pairs)
	if pairCount != (memberCount / 2) {
		t.Errorf("Expected %d; Actual %d\nMembers: %v\nPairs: %v\n", memberCount/2, pairCount, members, pairs)
	}
}
