package config_test

import (
	"fmt"
	"testing"

	"github.com/RobotDisco/mattermost-bagel/config"
)

func TestSQLitePersistenceConfig(t *testing.T) {
	fmt.Println("TestSQLitePersistenceConfig") // TODO: (TL) Doesn't find environment variables
	config.CreatePersistenceConfig()
	// TODO: (GCD) Write tests to cover both 'none' and 'sqlite' test cases.
	// TODO: (TL) Test DB with a test database file
}
