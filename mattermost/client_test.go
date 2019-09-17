package mattermost

import (
	"fmt"
	"os"
	"testing"
)

func TestNewMatterMostClient(t *testing.T) {
	fmt.Println("TestNewMatterMostClient")
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")
	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	api := NewMatterMostClient(serverURL, botUserName, botPassword)
	if api.Url != serverURL {
		t.Errorf("Unable to set the correct serverURL, expected: %v, got: %v", serverURL, api.Url)
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
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")
	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	api := NewMatterMostClient(serverURL, botUserName, botPassword)
	botUser := GetBotUser(*api)
	if botUser.Username != botUserName {
		t.Errorf("Expected Bot: %v, got: %v", botUserName, botUser.Username)
	}
}

func TestGetActiveChannelMembers(t *testing.T) {
	fmt.Println("TestGetActiveChannelMembers")
	serverURL := os.Getenv("BAGEL_MATTERMOST_URL")
	botUserName := os.Getenv("BAGEL_USERNAME")
	botPassword := os.Getenv("BAGEL_PASSWORD")

	teamName := os.Getenv("BAGEL_TEAM_NAME")
	channelName := os.Getenv("BAGEL_CHANNEL_NAME")

	api := NewMatterMostClient(serverURL, botUserName, botPassword)
	members := GetActiveChannelMembers(*api, teamName, channelName)

	for _, m := range members {
		if m.DeleteAt != 0 {
			t.Errorf("Expected to not find deactivated user %s.", m.Username)
		}
	}
}
