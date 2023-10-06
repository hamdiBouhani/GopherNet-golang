package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/hamdiBouhani/GopherNet-golang/dto"
	"github.com/hamdiBouhani/GopherNet-golang/storage"
	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
)

type BurrowService struct {
	storage storage.Storage
}

func NewBurrowService(storage storage.Storage) *BurrowService {
	return &BurrowService{storage: storage}
}

func InitialStates() ([]*model.Burrow, error) {

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get the current working directory: %v\n", err)
		return nil, err
	}

	dataDir := strings.TrimSuffix(currentDir, "services")
	// Create the absolute path to the JSON file
	jsonFilePath := filepath.Join(dataDir, "data", "initial.json")

	// Open our jsonFile
	jsonFile, err := os.Open(jsonFilePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("Successfully Opened initial.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := io.ReadAll(jsonFile)

	// Unmarshal the JSON data into a slice of Burrow structs
	var burrows []dto.BurrowDto
	if err := json.Unmarshal(byteValue, &burrows); err != nil {
		//log.Fatalf("Failed to unmarshal JSON: %v", err)
		return nil, err

	}

	// Print the burrows' details
	var res []*model.Burrow
	for _, burrow := range burrows {
		fmt.Printf("Burrow name: %s\n", burrow.Name)
		res = append(res, burrow.ParseToModel())
	}

	return res, nil
}

func (svc *BurrowService) InitialBurrowStates(burrow *model.Burrow) error {

	burrows, err := InitialStates()
	if err != nil {
		return err
	}

	if err := svc.storage.CreateManyBurrow(burrows); err != nil {
		return err
	}
	return nil
}
