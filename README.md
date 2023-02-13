## Requirements
- mysql 
- Go (version >= 1.18)

## Environment variables
- PORT: The port the server will run on (defaults to 8000)
- MYSQL_USER: username for MYSQL
- MYSQL_PASS: password for MYSQL
- MYSQL_URL: url and port for MYSQL E.G. 127.0.0.1:3306
## Setup
1. `git clone` the repository
2. `go get -u ./...` to get all dependencies
3. `go run *.go` to spin up a development server

`Gow` is recommended to watch the folder and rebuild on save 

## Database setup
db.go automatically creates the tables as needed