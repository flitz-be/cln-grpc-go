package main

import (
	"context"
	"fmt"

	"github.com/flitz-be/cln-grpc-go/cln"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	client := cln.NewNodeClient(conn)
	info, err := client.Getinfo(context.Background(), &cln.GetinfoRequest{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(info.Alias)
}
