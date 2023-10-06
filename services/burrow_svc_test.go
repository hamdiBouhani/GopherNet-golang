package services

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
)

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
