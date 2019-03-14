build:
	protoc --proto_path=./grpc --go_out=plugins=grpc:grpc ./grpc/rpc.proto
