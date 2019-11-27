package config_test

import (
	"fmt"
	"testing"

	"github.com/RobotDisco/mattermost-bagel/config"
)

func TestSQLitePersistenceConfig(t *testing.T) {
	fmt.Println("TestSQLitePersistenceConfig") // TODO: (TL) Doesn't find environment variables
	persistenceConfig := config.CreatePersistenceConfig()
	if persistenceConfig.PersistenceType == config.PersistenceNone {
		t.Errorf("SQLite was not configured!")
	}
	// TODO: (TL) Test DB with a test database file
}
