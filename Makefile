server: cmd/api/main.go
	go build -o bin/server cmd/api/main.go

genart: cmd/cli/main.go
	go build -o bin/genart cmd/cli/main.go

runserver: server
	./bin/server

clean:
	rm -rf bin
	go clean

.PHONY: server genart runserver clean
