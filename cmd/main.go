package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {

	netListen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen %v", err.Error())
	}

	grpcServer := grpc.NewServer()

	log.Printf("Sever started at %v", netListen.Addr())
	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to listend %v", err.Error())
	}
}
