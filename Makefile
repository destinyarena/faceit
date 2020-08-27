.PHONY: all build clean run docker-build docker-push docker proto-build

GORUN = go run
GOBUILD = go build

all: clean build

proto-build:
	rm -rf proto/*.go
	protoc --go_out=plugins=grpc:./proto -I ./proto ./proto/faceit.proto

clean:
	rm -rf bin

build: proto clean
	$(GOBUILD) -o bin/faceit cmd/faceit/*.go

docker-build:
	test $(DOCKERREPO)
	docker build . -t $(DOCKERREPO)

docker-push:
	test $(DOCKERREPO)
	docker push $(DOCKERREPO)

docker: docker-build docker-push
