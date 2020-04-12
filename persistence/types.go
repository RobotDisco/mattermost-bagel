package persistence


import (
	"os"
)

// BagelConfig contains the environment variables needed
// to connect to a mattermost server and team
// as a specified user, in a given channel
//
// Expected environment variables:
// - BAGEL_MATTERMOST_URL
// - BAGEL_USERNAME
// - BAGEL_PASSWORD
// - BAGEL_TEAM_NAME
// - BAGEL_CHANNEL_NAME
type BagelConfig struct {
	ServerURL   string
	BotUserName string
	BotPassword string
	TeamName    string
	ChannelName string
}

// CreateBagelConfig loads the required environment variables and creates
// a BagelConfig object encapsulating those values
func CreateBagelConfig() BagelConfig {
	return BagelConfig{
		ServerURL:   os.Getenv("BAGEL_MATTERMOST_URL"),
		BotUserName: os.Getenv("BAGEL_USERNAME"),
		BotPassword: os.Getenv("BAGEL_PASSWORD"),
		TeamName:    os.Getenv("BAGEL_TEAM_NAME"),
		ChannelName: os.Getenv("BAGEL_CHANNEL_NAME"),
	}
}
