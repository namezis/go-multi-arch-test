NAME := go-multi-arch-test

all: clean build

run:
	go run .

build:
	@echo "Version: " $(VERSION)

	echo "Build for amd64..."
	GOOS=linux GOARCH=amd64 go build -o $(NAME)-amd64 main.go

	echo "Build for armv7..."
	GOOS=linux GOARCH=arm GOARM=7 go build -o $(NAME)-armhf main.go

test:
	go test -v ./...

clean:
	-rm -fr $(NAME)*

lint:
	golangci-lint run -v

.PHONY: run build test clean
