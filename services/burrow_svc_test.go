package services

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/hamdiBouhani/GopherNet-golang/mocks"
	"github.com/hamdiBouhani/GopherNet-golang/storage/pg"
	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func MockBurrowService() *BurrowService {
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	service := &BurrowService{Storage: &pg.DBConn{Db: db}}
	err = service.Storage.Migrate()
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	return service
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func TestCurrentDir(t *testing.T) {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get the current working directory: %v\n", err)
		return
	}

	// Check if the string ends with "services"
	if strings.HasSuffix(currentDir, "services") {
		fmt.Println("The string ends with 'services'")
	} else {
		t.Error("The string does not end with 'services'")
		t.Fail()
		return
	}
}

func TestInitialStates(t *testing.T) {

	burrows, err := InitialStates()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(burrows) == 0 {
		t.Error("burrows is empty")
		t.Fail()
		return
	}
}

func TestInitialBurrowStates(t *testing.T) {

	svc := MockBurrowService()

	err := svc.InitialBurrowStates()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	res, err := svc.Storage.IndexBurrow() // Index the burrows
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if len(res) != 5 {
		t.Error("burrows should have 5 elements")
		t.Fail()
		return
	}

	err = svc.Storage.Drop()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestRentBurrow(t *testing.T) {
	svc := MockBurrowService()

	newBurrow := mocks.MockBurrow(false)

	err := svc.Storage.CreateBurrow(newBurrow)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	err = svc.RentBurrow(newBurrow.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	updatedBurrow, err := svc.Storage.ShowBurrow(newBurrow.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if !updatedBurrow.Occupied {
		t.Error("burrow should be occupied")
		t.Fail()
		return
	}

	err = svc.Storage.Drop()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
}

func TestRunUpdateStatusTask(t *testing.T) {
	svc := MockBurrowService()

	newBurrow := mocks.MockBurrow(false)
	err := svc.Storage.CreateBurrow(newBurrow)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	err = svc.RentBurrow(newBurrow.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	burrow1, err := svc.Storage.ShowBurrow(newBurrow.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if !burrow1.Occupied {
		t.Error("burrow should be occupied")
		t.Fail()
		return
	}

	sleepDuration := 10 * time.Second
	err = svc.RunUpdateStatusTask(10 * time.Second)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	fmt.Println("Sleeping for", sleepDuration)
	time.Sleep(sleepDuration)
	fmt.Println("Action performed at", time.Now())
	burrow2, err := svc.Storage.ShowBurrow(newBurrow.ID)
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

	if burrow2.Depth < burrow1.Depth {
		t.Error("burrow depth should be less than before")
		t.Fail()
		return
	}

	err = svc.Storage.Drop()
	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}

}
