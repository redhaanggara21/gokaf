Compile Protoc Unary / Streaming gRPC

1. protoc --go_out=plugins=grpc:. **/*.proto
2. go run cmd/main.go
3. protoc *.proto --go_out=.
