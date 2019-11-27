package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RobotDisco/mattermost-bagel/mattermost"
	_ "github.com/mattn/go-sqlite3" // Add SQLite support
)

const (
	// DefaultStaleBoundary is fallback value if the BAGEL_SQLITE_STALE_BOUNDARY environment variable is empty (default is 35 days)
	DefaultStaleBoundary = "35 days"
)

// CreatePersistenceConfig ... TODO: (TL) Comment
func CreatePersistenceConfig() PersistenceConfig {
	if strings.ToLower(os.Getenv("BAGEL_PERSISTENCE_METHOD")) != "sqlite" {
		return PersistenceConfig{}
	}

	// SQLite DB File
	sqliteFile := os.Getenv("BAGEL_SQLITE_FILE")
	if sqliteFile == "" {
		// TODO: (TL) Fall back to a DB name?
		fmt.Printf("'BAGEL_PERSISTENCE_METHOD' of 'sqlite' specified, but 'BAGEL_SQLITE_FILE' is not specified.\nFalling back to no persistence model.\n")
		return PersistenceConfig{}
	}

	// SQLite DB
	database, err := sql.Open("sqlite3", sqliteFile)
	if err != nil {
		fmt.Printf("Encountered error connecting to SQLite DB: %+v\nFalling back to 'PersistenceMethodNone'.\n", err)
		return PersistenceConfig{}
	}

	// SQLite Create Table - see client.go for a sample GetChannelMembers call result TL;DR: ChannelId is a hash-type string, UserId is a hash-type string
	// the 'pairing' column will be a concatenation of two `UserId` fields delimited by a colon, i.e. "3eizzrqatin3te8s3a3yj5r6xc:3eizzrqatin3te8s3a3yj5r6xd"
	createTableStatement, err := database.Prepare("CREATE TABLE IF NOT EXISTS scheduled_outings (message_date INTEGER DEFAULT CURRENT_TIMESTAMP, channel_id TEXT, pairing TEXT)")
	if err != nil {
		fmt.Printf("Encountered error preparing CREATE TABLE statement: %+v\nFalling back to 'PersistenceMethodNone'.\n", err)
		return PersistenceConfig{}
	}
	createTableStatement.Exec()

	// SQLite INSERT statement
	insertMeetupStatement, err := database.Prepare("INSERT INTO scheduled_outings (channel_id, pairing) VALUES (?, ?)")
	if err != nil {
		fmt.Printf("Encountered error preparing INSERT statement: %+v\nFalling back to 'PersistenceMethodNone'.\n", err)
		return PersistenceConfig{}
	}

	// SQLite SELECT statement
	timeframe := os.Getenv("BAGEL_SQLITE_STALE_BOUNDARY")
	if timeframe == "" {
		// Fall back to default 5 weeks
		fmt.Printf("No value specified in 'BAGEL_SQLITE_STALE_BOUNDARY', falling back to %s.\n", DefaultStaleBoundary)
		timeframe = DefaultStaleBoundary
	}
	// Find all outings that were previously scheduled between now and <timeframe> for the given channel and pair of users
	selectMeetupsStatement := fmt.Sprintf("SELECT COUNT(1) FROM scheduled_outings WHERE message_date > datetime('now', '-%s') AND pairing = ?", timeframe)

	retryCount := 1
	retryEnvironmentVar := os.Getenv("BAGEL_PERSISTENCE_RETRY_COUNT")
	if retryEnvironmentVar != "" {
		retryCount, _ = strconv.Atoi(retryEnvironmentVar)
	}

	return PersistenceConfig{
		PersistenceSQLite,
		retryCount,
		SQLiteClient{
			database,
			insertMeetupStatement,
			selectMeetupsStatement,
		},
	}
}

// VerifyPairs uses the active PersistenceConfig to determine if matches have been previously made.
// When using the PersistenceNone option, this is a simple passthrough.
func (persistenceConfig PersistenceConfig) VerifyPairs(pairs mattermost.ChannelMemberPairs) VerifyResult {
	if persistenceConfig.PersistenceType == PersistenceNone {
		return VerifyResult{
			Successes: pairs,
		}
	}
	var successes mattermost.ChannelMemberPairs
	var failures mattermost.ChannelMemberPairs
	for _, pair := range pairs {
		pairIdentifier := pair.Identifier()
		fmt.Print(persistenceConfig.sqliteClient.selectMeetupsStatement)
		fmt.Printf(" [%s]\n", pairIdentifier)
		rows, err := persistenceConfig.sqliteClient.database.Query(persistenceConfig.sqliteClient.selectMeetupsStatement, pairIdentifier)
		defer rows.Close()
		if err != nil {
			fmt.Printf("Encountered error with SELECT query: %+v, assuming failure.\n", err)
			failures = append(failures, pair)
		} else if rows.Next() { // Only expect one row
			var previousPairingsWithinDateRange int
			rows.Scan(&previousPairingsWithinDateRange)
			if previousPairingsWithinDateRange == 0 { // Success, no records found!
				fmt.Printf("No record exists, marking as a successful pair: %s\n", pair.Identifier())
				successes = append(successes, pair)
			} else { // Failure, try again!
				fmt.Printf("A record exists for this pair: %s, storing as a leftover.\n", pair.Identifier())
				failures = append(failures, pair)
			}
		} else {
			fmt.Printf("A record exists for this pair: %s, storing as a leftover (%+v).\n", pair.Identifier(), err)
			failures = append(failures, pair)
		}
	}
	return VerifyResult{
		successes,
		failures,
	}
}

// LogPairs considers the PersistenceType and outputs the pairs as appropriate (console, SQLite, etc.)
func (persistenceConfig PersistenceConfig) LogPairs(pairs mattermost.ChannelMemberPairs) {
	switch persistenceConfig.PersistenceType {
	case PersistenceNone:
		persistenceConfig.logPairsToConsole(pairs)
	case PersistenceSQLite:
		persistenceConfig.logPairsToSQLite(pairs)
	}
}

func (persistenceConfig PersistenceConfig) logPairsToConsole(pairs mattermost.ChannelMemberPairs) {
	for index, pair := range pairs {
		pairIdentifier := pair.Identifier()
		fmt.Printf("Pair #%d: %s\n", index, pairIdentifier)
	}

}

func (persistenceConfig PersistenceConfig) logPairsToSQLite(pairs mattermost.ChannelMemberPairs) {
	for _, pair := range pairs {
		pairIdentifier := pair.Identifier()
		_, err := persistenceConfig.sqliteClient.insertMeetupStatement.Exec(pair.ChannelID, pairIdentifier)
		if err != nil {
			fmt.Printf("Encountered error saving pair: %s (%+v).\n", pair.Identifier(), err)
		}
	}

}
