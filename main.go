package main

import (
	"log"
	"net"

	"github.com/ybalexdp/sample-grpc/pb"
	"github.com/ybalexdp/sample-grpc/service"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":2007")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	sampleService := &service.SampleService{}
	pb.RegisterSampleServer(server, sampleService)
	server.Serve(listenPort)
}
