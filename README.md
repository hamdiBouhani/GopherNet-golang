# GopherNet-golang
a platform for managing rentals of gopher burrows.



## Installation
```bash
make postgres
make createdb
```
## Usage
```bash
    2023/10/07 10:58:51 Using Postgres Database
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
or install jq and run:
```bash
sudo chmod +x ./curl_burrow_status.sh 
./curl_burrow_status.sh
```
