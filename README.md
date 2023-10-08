# GopherNet-golang
a platform for managing rentals of gopher burrows.

## Project Structure
```bash
.
├── app_architecture.drawio.png
├── architecture-img.png
├── Architecture_project.md
├── Architecture_project.pdf
├── cloud_architecture.drawio.png
├── cloud_architecture.png
├── cmd
├── cover.out
├── curl_burrow_status.sh
├── curl_rent_burrow.sh
├── data
│   └── initial.json
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── dto
│   ├── burrow.go
│   └── http_responses.go
├── gain.pro senior go engineer assignment.pdf
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
├── mocks
│   └── fakes.go
├── README.md
├── report.txt
├── server
│   ├── gorm.db
│   ├── helper_func.go
│   ├── http_burrow.go
│   ├── http_burrow_test.go
│   └── http.go
├── services
│   ├── burrow_svc.go
│   ├── burrow_svc_test.go
│   └── gorm.db
├── storage
│   ├── model
│   │   └── burrow.go
│   ├── pg
│   │   ├── burrow.go
│   │   ├── db.go
│   │   └── db_test.go
│   └── storage.go
└── utils
    └── helper_func.go
```



## Installation
```bash
make postgres
make createdb
```
## Usage
```bash
    Usage:
    gopherne [flags]
    gopherne [command]

    Available Commands:
    completion   Generate the autocompletion script for the specified shell
    help         Help about any command
    load-burrows Load Burrows from initial.json
    rest         Run rest server

    Flags:
    -h, --help   help for gopherne

    Use "gopherne [command] --help" for more information about a command.
```

if you want to load the initial burrows:
```bash
go run main.go load-burrows
```

to test the rest server:
```bash
go run main.go rest
```

then in another terminal:
```bash
curl -X GET http://localhost:8080/burrows
```

or install jq:
```bash
# For Linux
sudo apt-get update && sudo apt-get install jq
# For Mac
brew install jq
# For Windows with chocolatey
chocolatey install jq
```
and run:
```bash
sudo chmod +x ./curl_burrow_status.sh 
./curl_burrow_status.sh
```


```bash
sudo chmod +x ./curl_rent_burrow.sh 
./curl_rent_burrow.sh 9

9 is the burrow id
```

---

1. **Install golangci-lint**:

```bash
# binary will be $(go env GOPATH)/bin/golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2

golangci-lint --version
```

2. **Run linter**:

```bash
make run-lint
```

--- 

## Solutions

**Requirements 1:**

I created command line application to load the initial burrows from `data/initial.json` file.

Load the initial burrows:
```bash
go run main.go load-burrows
```

**Requirements 2:**

I used golang [Ticker](https://gobyexample.com/tickers) to run background task to Update the burrows status every 5 seconds.

```go
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
					err := svc.Storage.UpdateBurrowAttributes(b.ID, map[string]interface{}{"occupied": false, "age": b.Age, "depth": b.Depth})
					if err != nil {
						log.Printf("Error updating burrow: %v\n", err)
					}

				} else {
					err := svc.Storage.UpdateBurrowAttributes(b.ID, map[string]interface{}{"age": b.Age, "depth": b.Depth})
					if err != nil {
						log.Printf("Error updating burrow: %v\n", err)
					}
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

```


**Requirements 3 & 4:**

I created a rest server to manage the burrows.  

to run the rest server:
```bash
go run main.go rest
```

**Requirements 5:**

I used golang [Ticker](https://gobyexample.com/tickers) to run background task to Create a report  summarising the total depth and number of available burrows.

```go

// Reporting: Create a report summarising the total depth and number of available
// burrows, as well as the largest and smallest burrows - by volume.
// You can assume the burrows are cylindrical (where width is the diameter).
// This reportshould be outputted as a plain text file every 10 minutes.
func (svc *BurrowService) Report(duration time.Duration) error {

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
	ticker := time.NewTicker(duration)
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
			fmt.Printf("Burrow name: %s, Volume: %.2f\n", burrow.Name, volume)
			if volume >= largestVolume {
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
		file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file: %v\n", err)
			return err
		}

		// Write the report to the file
		if _, err := file.WriteString(report); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
		}

		// Write the report to the file
		if _, err := file.WriteString("--------------------------------\n"); err != nil {
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

```
