
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro:v2.9.3 -I ./ --micro_out=./ --go_out=./ --go_out=./ .proto/gateway/gateway.proto

	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/gateway/gateway.proto
    

.PHONY: build
build: proto

	go build -o gateway-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t gateway-service:latest
