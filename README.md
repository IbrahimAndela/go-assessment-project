# Go Rest-Api server

## Requirements
```
gorm: go get -u github.com/jinzhu/gorm
mux: go get -u github.com/gorilla/mux
```

## Database
Postgres

## Flags
- host=Server host
- port=Server port
- user=Database username
- dbname=Database name
- password=Database password

## Commands
To run server
`` go run main.go -host=hostname -user=assess1 -dbname=assessment -password=1234test``
