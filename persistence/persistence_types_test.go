package persistence

import (
	"fmt"
	"testing"
)

func TestSQLitePersistenceConfig(t *testing.T) {
	fmt.Println("TestSQLitePersistenceConfig") // TODO: (TL) Doesn't find environment variables
	CreatePersistenceConfig()
	// TODO: (GCD) Write tests to cover both 'none' and 'sqlite' test cases.
	// TODO: (TL) Test DB with a test database file
}
