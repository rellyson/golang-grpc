package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

func Bootstrap() (*grpc.Server, error) {
	adr := os.Getenv("LISTEN_ADR")
	lis, err := net.Listen("tcp", adr)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return nil, err
	}

	g := grpc.NewServer()
	SetHandlers(g)

	log.Print("Ready to listen in: ", lis.Addr())

	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
		return nil, err
	}

	return g, nil
}
