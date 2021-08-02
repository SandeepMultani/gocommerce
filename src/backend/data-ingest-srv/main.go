package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	event "github.com/SandeepMultani/gocommerce/src/backend/data-ingest-srv/protobuf/event"
	"github.com/SandeepMultani/gocommerce/src/backend/data-ingest-srv/server"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 5501, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	ingestServer := server.NewIngestServer()
	event.RegisterIngestServiceServer(grpcServer, ingestServer)
	grpcServer.Serve(lis)
}
