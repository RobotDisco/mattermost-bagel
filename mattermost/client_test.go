package mattermost

import (
	"fmt"
	"testing"

	"github.com/RobotDisco/mattermost-bagel/config"
)

var cfg = config.Config{
 	MattermostURL: "http://localhost:8065",
 	MattermostUser: "coffeebot1",
 	MattermostPassword: "password",
 	MattermostTeam: "test",
 	MattermostChannel: "coffeebot",
}

func TestNewMatterMostClient(t *testing.T) {
	fmt.Println("TestNewMatterMostClient")

	api := NewMatterMostClient(cfg.MattermostURL, cfg.MattermostUser, cfg.MattermostPassword)
	if api.Url != cfg.MattermostURL {
		t.Errorf("Unable to set the correct serverURL, expected: %v, got: %v", cfg.MattermostURL, api.Url)
	}
	/* Not working? Login doesn't seem to happen
	if api.AuthToken == "" {
		t.Errorf("Unable to establish a proper AuthToken for User: %v", botUserName)
	}
	if api.AuthType == "" {
		t.Errorf("AuthType is empty for User: %v", botUserName)
	}
	*/
}

func TestGetBotUser(t *testing.T) {
	fmt.Println("TestGetBotUser")

	api := NewMatterMostClient(cfg.MattermostURL, cfg.MattermostUser, cfg.MattermostPassword)
	botUser := GetBotUser(*api)
	if botUser.Username != cfg.MattermostUser {
		t.Errorf("Expected Bot: %v, got: %v", cfg.MattermostUser, botUser.Username)
	}
}

func TestGetActiveChannelMembers(t *testing.T) {
	fmt.Println("TestGetActiveChannelMembers")

	api := NewMatterMostClient(cfg.MattermostURL, cfg.MattermostUser, cfg.MattermostPassword)
	_, members := GetActiveChannelMembers(*api, cfg.MattermostTeam, cfg.MattermostChannel)

	for _, m := range members {
		if m.DeleteAt != 0 {
			t.Errorf("Expected to not find deactivated user %s.", m.Username)
		}
	}
}
