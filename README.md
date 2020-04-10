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
