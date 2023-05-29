CONFIG=cmd/atm-controller/config/viper/example/config.yaml

all: run

run-now:
	./bin/atm-controller --config $(CONFIG)

run: build run-now

run-force: build-force run-now

build: test bin/atm-controller

build-force: bin/atm-controller

test:
	go test -v ./... -cover

bin/atm-controller: bin cmd/atm-controller
	go build -o bin/atm-controller cmd/atm-controller/main.go

bin:
	mkdir bin
