build:
	protoc -I . server.proto --go_out=plugins=grpc:serveur
