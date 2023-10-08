package pg

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func TestCreateConnection(t *testing.T) {
	t.Skip()
	db := NewDBConn()
	defer db.Close()

	err := db.CreateConnection()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if db.Db == nil {
		t.Error("db is nil")
		t.Fail()
		return
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.Db.DB()
	if err != nil {
		log.Printf("Error updating burrow: %v\n", err)
		t.Error(err)
		t.Fail()
		return
	}

	// Ping
	err = sqlDB.Ping()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	t.Logf("Connected to database %v", sqlDB.Stats())
}
