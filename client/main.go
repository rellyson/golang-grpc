package main

import (
	"context"
	"log"
	"os"

	e "github.com/rellyson/golang-grpc/proto/gen-code/examples"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	adr := os.Getenv("LISTEN_ADR")
	conn, err := grpc.Dial(adr, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := e.NewExampleServiceClient(conn)

	res, err := c.SayHello(context.Background(), &e.HelloRequest{Name: "Ednaldo Pereira"})

	if err != nil {
		log.Fatalf("Error calling server: %s", err)
	}

	log.Printf("Response: %s", res.Message)
}
