# Golang bindings for the CLN gRPC client

```
protoc --go_out=cln --go_opt=paths=source_relative \
    --go-grpc_out=cln --go-grpc_opt=paths=source_relative primitives.proto
protoc --go_out=cln --go_opt=paths=source_relative \
    --go-grpc_out=cln --go-grpc_opt=paths=source_relative node.proto
go mod tidy
```

An example client can be found in the `example` folder.