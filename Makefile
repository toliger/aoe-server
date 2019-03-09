build:
	protoc --proto_path=. --go_out=plugins=grpc:serveur rpc.proto
