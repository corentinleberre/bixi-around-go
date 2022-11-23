.DEFAULT_GOAL := help
SHELL := /bin/bash

#help: @ list available tasks on this project
help:
	@grep -E '[a-zA-Z\.\-]+:.*?@ .*$$' $(MAKEFILE_LIST)| tr -d '#'  | awk 'BEGIN {FS = ":.*?@ "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

#build: @ build bixi-around binary
build:
	@echo "[BUILD] building bixi-around binary"
	rm -rf ./bin
	go build -o ./bin/bixi-around ./src

#run: @ launch bixi-around binary
run: build
	@echo "[RUN] Launching bixi-around-go"
	./bin/bixi-around

#dev: @ launch bixi-around in dev mode
dev:
	@echo "[RUN] Launching bixi-around-go"
	go run ./src