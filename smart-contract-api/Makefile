all: build

VERSION?=latest

build:
	go build -o bin/smart-contract-api .

run:
	go run main.go

docker-build:
	docker build -f docker/Dockerfile -t gabilb/smart-contract-api:$(VERSION) .

docker-push:
	docker push gabilb/smart-contract-api:$(VERSION)
