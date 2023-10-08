# GopherNet-golang
a platform for managing rentals of gopher burrows.



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


1 - **Requirements 1:**

I created command line application to load the initial burrows from `data/initial.json` file.

Load the initial burrows:
```bash
go run main.go load-burrows
```

2 - **Requirements 2:**

I used golang [Ticker](https://gobyexample.com/tickers) to run background task to Update the burrows status every 5 seconds.


3/4 - **Requirements 3 / 4:**

I created a rest server to manage the burrows.  

to run the rest server:
```bash
go run main.go rest
```

5 - **Requirements r:**

I used golang [Ticker](https://gobyexample.com/tickers) to run background task to Create a report  summarising the total depth and number of available burrows.
