package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/flitz-be/cln-grpc-go/cln"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func loadTLSCredentials(caCertPath, clientCertPath, clientKeyPath string) (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func main() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	credentials, err := loadTLSCredentials(
		fmt.Sprintf("%s/.lightning/ca.pem", dirname),
		fmt.Sprintf("%s/.lightning/client.pem", dirname),
		fmt.Sprintf("%s/.lightning/client-key.pem", dirname))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(credentials),
	}

	conn, err := grpc.Dial("localhost:9737", opts...)
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
