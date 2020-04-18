## Minesweeper API
------------------------------------------------------------

This is a basic API for the game Minesweeper. Written in Golang and using Redis for storage.

For build it locally, you might need Redis installed locally, otherwise, use docker-compose to run via make command,
or directly run with *docker-compose*


## Using golang and remote Redis.

### Run the tests
```shell script
go test -v ./...
```

### Build the API
```shell script
go build -o ms-api cmd/main.go
```

### Run the API
```shell script
./ms-api
```

## With Docker compose
### up all the containers, run locally in 8080 port
```shell script
make up
```
or
```shell script
docker-compose up
```

### down all the containers
```shell script
make down
```
or
```shell script
docker-compose down --remove-orphans
```

### check status
```shell script
docker-compose ps
```

### Create a new user (Server response: 201)
```shell script
curl -X POST \
  http://localhost:8080/users \
  -d '{
	"username": "tomasito"
}'
```

### Create a game (need a username already created, Server response: 201 otherwise 400 will be sent to the client)
```shell script
curl -X POST \
  http://localhost:8080/games \
  -d '{
	"name": "game1",
	"rows": 10,
	"cols": 10,
	"mines": 10,
	"username": "tomasito"
}'
```

### Start a game (Server response: 200, if the game or username are not present, 400 will be sent to the client)
```shell script
localhost:8080/games/game1/users/tomasito
```

### Play clicking, marking or flagging (click_type might be click, mark or flag, Server response: 200)
```shell script
curl -X POST \
  http://localhost:8080/games/game1/users/tomasito/click \
  -d '{
	"row": 1,
	"col": 3,
	"click_type": "click"
}'
```

### Demo URL
```
http://ms-tomas-api.herokuapp.com/heartbeat
```

### Check out the CLI for this minesweeper
```
https://github.com/nicolasacquaviva/minesweeper-cli
```
