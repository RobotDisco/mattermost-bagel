package config

import (
	"os"
)

type Config struct {
	MattermostURL string
	MattermostUser string
	MattermostPassword string
	MattermostTeam string
	MattermostChannel string

	DatabaseBackend string
	NonRepeatingPairsRetryAttempts string

	SQLiteFile string
	SQLiteStaleBoundary string
}

func CreateConfigFromEnvironmentVariables() Config {
	return Config{
		MattermostURL: os.Getenv("BAGEL_MATTERMOST_URL"),
		MattermostUser: os.Getenv("BAGEL_USERNAME"),
		MattermostPassword: os.Getenv("BAGEL_PASSWORD"),
		MattermostTeam: os.Getenv("BAGEL_TEAM_NAME"),
		MattermostChannel: os.Getenv("BAGEL_CHANNEL_NAME"),

		DatabaseBackend: os.Getenv("BAGEL_DB_BACKEND"),
		NonRepeatingPairsRetryAttempts: os.Getenv("BAGEL_NONREPEATING_PAIRS_MAX_RETRY_ATTEMPTS"),

		SQLiteFile: os.Getenv("BAGEL_SQLITE_FILE"),
		SQLiteStaleBoundary: os.Getenv("BAGEL_SQLITE_STALE_BOUNDARY"),
	}
}
