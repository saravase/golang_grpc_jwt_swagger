gen:
	protoc --proto_path=proto/ proto/*.proto --go_out=plugins=grpc:pb
clean:
	rm pb/*.go
ps:
	go run server/main.go
pc:
	go run client/main.go