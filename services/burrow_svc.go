package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/hamdiBouhani/GopherNet-golang/dto"
	"github.com/hamdiBouhani/GopherNet-golang/storage"
	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
	"github.com/hamdiBouhani/GopherNet-golang/utils"
)

type BurrowService struct {
	Storage storage.Storage
}

func NewBurrowService(storage storage.Storage) *BurrowService {
	return &BurrowService{Storage: storage}
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

func (svc *BurrowService) InitialBurrowStates() error {

	burrows, err := InitialStates()
	if err != nil {
		return err
	}

	err = svc.Storage.CreateManyBurrow(burrows)
	if err != nil {
		return err
	}
	return nil
}

func (svc *BurrowService) UpdateBurrow(id int64) error {

	burrow, err := svc.Storage.ShowBurrow(id)
	if err != nil {
		return err
	}

	if err := svc.Storage.SaveBurrow(burrow); err != nil {
		return err
	}
	return nil

}

// Rent Burrow: Implement an HTTP REST endpoint to handle requests for rentinga burrow.
// If a burrow is available (not occupied and hasn't collapsed), the burrowwill be rented and it’s status will be updated.
// Otherwise, return an appropriateerror message.
func (svc *BurrowService) RentBurrow(id int64) error {

	burrow, err := svc.Storage.ShowBurrow(id)
	if err != nil {
		return err
	}

	if burrow.Occupied || burrow.Age > 25 {
		return fmt.Errorf("burrow is not available")
	}

	burrow.Occupied = true

	if err := svc.Storage.SaveBurrow(burrow); err != nil {
		return err
	}
	return nil

}

// Burrow Status: Provide a REST endpoint to fetch the current status of theburrows
func (svc *BurrowService) BurrowStatus() ([]*model.Burrow, error) {
	res, err := svc.Storage.IndexBurrow()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (svc *BurrowService) RunUpdateStatusTask(duration time.Duration) error {
	burrows, err := svc.Storage.IndexBurrow() // Index the burrows
	if err != nil {
		return err
	}

	if len(burrows) == 0 {
		return fmt.Errorf("no burrows found")
	}

	for _, burrow := range burrows {
		b := burrow
		go func() {
			job := func() {
				fmt.Println("Running the job at", time.Now())
				fmt.Printf("Burrow name: %s\n", b.Name)
				b.Age++
				fmt.Printf("Burrow name: %s age increased to %d at % v \n", b.Name, b.Age, time.Now())

				if b.Occupied {
					b.Depth += (b.Depth * 0.9)
				}

				if b.Occupied && ((b.Age / 1440) >= 25) { //Burrow age (A, in minutes), with each burrow lasting exactly 25 days before collapsing.
					b.Occupied = false
					svc.Storage.UpdateBurrowAttributes(b.ID, map[string]interface{}{"occupied": false, "age": b.Age, "depth": b.Depth})
				} else {
					svc.Storage.UpdateBurrowAttributes(b.ID, map[string]interface{}{"age": b.Age, "depth": b.Depth})
				}
				log.Printf("Burrow state: %+v\n", b)
			}

			// Create a ticker that triggers every minute
			ticker := time.NewTicker(duration)

			// Run the job when the ticker triggers
			for range ticker.C {
				job() // Execute the job
			}
		}()
	}

	return nil
}

// Reporting: Create a report summarising the total depth and number of available
// burrows, as well as the largest and smallest burrows - by volume.
// You can assume the burrows are cylindrical (where width is the diameter).
// This reportshould be outputted as a plain text file every 10 minutes.
func (svc *BurrowService) Report() error {

	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get the current working directory: %v\n", err)
		return err
	}

	dataDir := strings.TrimSuffix(currentDir, "services")
	// Create the absolute path to the JSON file
	fileName := filepath.Join(dataDir, "report.txt")

	// Create a ticker that triggers every minute
	ticker := time.NewTicker(10 * time.Minute)
	// Run the job when the ticker triggers
	for range ticker.C {

		burrows, err := svc.Storage.IndexBurrow()
		if err != nil {
			return err
		}

		if len(burrows) == 0 {
			return fmt.Errorf("Report : no burrows found")
		}

		// Calculate total depth, number of burrows, and find the largest and smallest burrows
		var totalDepth float64
		var largestVolume float64 = 0
		var smallestVolume float64 = math.MaxFloat64
		var largestBurrow *model.Burrow
		var smallestBurrow *model.Burrow

		for _, burrow := range burrows {
			totalDepth += burrow.Depth

			volume := utils.CalculateVolume(burrow)
			if volume > largestVolume {
				largestVolume = volume
				largestBurrow = burrow
			}
			if volume < smallestVolume {
				smallestVolume = volume
				smallestBurrow = burrow
			}
		}

		// Prepare the report
		report := fmt.Sprintf("Report (Generated at %s):\n", time.Now())
		report += fmt.Sprintf("Total Depth: %.2f\n", totalDepth)
		report += fmt.Sprintf("Number of Burrows: %d\n", len(burrows))
		report += fmt.Sprintf("Largest Burrow - Name: %s, Volume: %.2f\n", largestBurrow.Name, largestVolume)
		report += fmt.Sprintf("Smallest Burrow - Name: %s, Volume: %.2f\n", smallestBurrow.Name, smallestVolume)

		// Open the file in write mode
		file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return err
		}

		// Write the report to the file
		if _, err := file.WriteString(report); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}

		// Close the file
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %v\n", err)
		}

		fmt.Println("Report generated and saved to", fileName)
	}
	return nil
}
