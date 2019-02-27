build:
	protoc --proto_path=. --go_out=./rpc rpc.proto
