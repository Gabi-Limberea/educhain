all: build

VERSION?=latest

build:
	go build -o bin/ipfs-api .

run:
	go run main.go

docker-build:
	docker build -f docker/Dockerfile -t gabilb/ipfs-api:$(VERSION) .

docker-push:
	docker push gabilb/ipfs-api:$(VERSION)
