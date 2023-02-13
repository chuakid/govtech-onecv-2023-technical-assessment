## Requirements
- mysql 
- Go (version >= 1.16)

## Environment variables
- jwtkey: key for JWT generation, can be any random string
- PORT: The port the server will run on (defaults to 8000)
- MYSQL_USER
- MYSQL_PASS 
- MYSQL_URL: database connection string for postgresql
## Setup
1. `git clone` the repository
2. `go get -u ./...` to get all dependencies
3. `go *.go` to spin up a development server

`Gow` is recommended to watch the folder and rebuild on save 
