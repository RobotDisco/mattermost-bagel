package config

import (
	"database/sql"

	"github.com/RobotDisco/mattermost-bagel/mattermost"
)

// PersistenceConfig represents the environment variables
// used to configure the persistence portion of Bagel
//
// Expected environment variables:
// - BAGEL_DB_BACKEND (Options: "none", "sqlite")
// - BAGEL_NONREPEATING_PAIRS_MAX_RETRY_ATTEMPTS (Defaults to 0 for "none", 1 for everything else)
//
// SQLite Only
// - BAGEL_DB_SQLITE_FILE (Local path to the file)
type PersistenceConfig struct {
	PersistenceType PersistenceMethod
	RetryCount      int
	sqliteClient    SQLiteClient
}

// PersistenceMethod is an enum mapping the various persistence storage types available
type PersistenceMethod int

const (
	// PersistenceNone specifies that no persistence will be used
	PersistenceNone PersistenceMethod = 0
	// PersistenceSQLite specifies that a local SQLite file DB is used
	PersistenceSQLite PersistenceMethod = 1
)

// TODO: (TL) Add functions to the SQLiteClient so we're not manipulating raw stuffs

// SQLiteClient wraps up the various pieces needed to SQLite db manipulation
type SQLiteClient struct {
	database               *sql.DB
	insertMeetupStatement  *sql.Stmt
	selectMeetupsStatement string
}

// VerifyResult ... TODO: (TL) Comment
type VerifyResult struct {
	Successes mattermost.ChannelMemberPairs
	Failures  mattermost.ChannelMemberPairs
}
