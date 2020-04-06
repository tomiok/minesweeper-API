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