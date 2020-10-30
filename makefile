gen:
	protoc --proto_path=proto/ proto/*.proto --go_out=plugins=grpc:pb
clean:
	rm pb/*.go
ps:
	go run grpc/server/main.go
pc:
	go run grpc/client/main.go