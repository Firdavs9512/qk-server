# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_CLI=qk-manager-cli
BINARY_NAME_SERVER=qk-manager-server

build:
	$(GOBUILD) -o $(BINARY_NAME_CLI) -v cmd/cli/main.go
	$(GOBUILD) -o $(BINARY_NAME_SERVER) -v cmd/server/main.go

test:
	$(GOTEST) -v ./test/...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME_CLI)
	rm -f $(BINARY_NAME_SERVER)

cli:
	$(GORUN) cmd/cli/main.go

server:
	$(GORUN) cmd/server/main.go
