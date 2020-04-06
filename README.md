## Minesweeper API
------------------------------------------------------------

### Run the tests
```shell script
go test -v ./...
```

### Build the API
```shell script
go build -o ms ./cmd
```

### Run the API
```shell script
./ms
```

### Create a new user
```shell script
curl -X POST \
  http://localhost:8080/users \
  -d '{
	"username": "tomasito"
}'
```

### Create a game (need a username already created), otherwise 400 will be received
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

### Start a game
```shell script
localhost:8080/games/game1/users/tomasito
```