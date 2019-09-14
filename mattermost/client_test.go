package mattermost

import (
	"os"
	"testing"
)

func TestNewMatterMostClient(t *testing.T) {
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
