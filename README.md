# Golang bindings for the CLN gRPC client

Generating from [CLN's proto files](https://github.com/ElementsProject/lightning/tree/master/cln-grpc/proto):

```
protoc --go_out=cln --go_opt=paths=source_relative \
    --go-grpc_out=cln --go-grpc_opt=paths=source_relative primitives.proto
protoc --go_out=cln --go_opt=paths=source_relative \
    --go-grpc_out=cln --go-grpc_opt=paths=source_relative node.proto
go mod tidy
```

An example client can be found in the `example` folder. Make sure that your node has gRPC enabled
(`--grpc-port=port`), and change the host name in `example/main.go` or use port-forwarding.
To use the client, you need the following files from your node's `.lightning` directory:

- ca.pem  
- client-key.pem 
- client.pem