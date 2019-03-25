build:
	protoc --proto_path=./grpc --go_out=plugins=grpc:grpc ./grpc/rpc.proto

test:
	golint -set_exit_status ./... && go test -race -short ./... && errcheck ./... && go test -cover ./...

run:
	docker build -t aoe-server . && docker run -p 50010:50010 aoe-server
