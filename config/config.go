package config

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
