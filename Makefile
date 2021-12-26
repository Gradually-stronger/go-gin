.PHONY: start build build-linux swagger

NOW = $(shell date -u '+%Y%m%d%I%M%S')

SERVER_BIN = "./services/server"
RELEASE_ROOT = "release"
RELEASE_SERVER = "release/server"

all: start

build:
	@go build -ldflags "-w -s" -o $(SERVER_BIN) ./services

start:
	@go build -o $(SERVER_BIN) ./services
	$(SERVER_BIN) -c config/config.toml -m config/model.conf -swagger ./internal/app/swagger

swagger:
	swaggo -s internal/app/routers/web/swagger.go -p . -o internal/app/swagger

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o $(SERVER_BIN) ./sevices

publish: build-linux
	scp $(SERVER_BIN) root@39.98.250.155:/root/services/newparkassets/gxtparkassets2

publish-swagger: swagger
	scp ./internal/app/swagger/* root@39.98.250.155:/root/services/newparkassets/swagger/

test:
	@go test -cover -race ./...

clean:
	rm -rf data release $(SERVER_BIN) ./internal/app/test/data ./services/data

pack: build
	rm -rf $(RELEASE_ROOT)
	mkdir -p $(RELEASE_SERVER)
	cp -r $(SERVER_BIN) configs $(RELEASE_SERVER)
	cd $(RELEASE_ROOT) && zip -r server.$(NOW).zip "server"
