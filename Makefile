REPOSITORY=perriea/webserver
VERSION=0.0.1

all: build

build:
	@echo "Build binary"
	@mkdir -p ./build
	@go build -ldflags="-w -s" -o ./build/webserver .

container:
	@docker build -t ${REPOSITORY}:${VERSION} --no-cache .

clean:
	@rm -rf build/

.PHONY: help build container clean