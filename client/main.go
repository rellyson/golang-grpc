package main

import (
	"context"
	"log"
	"os"

	cpb "github.com/rellyson/golang-grpc/proto/gen-code"
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

	c := cpb.NewCurrencyServiceClient(conn)

	res, err := c.ConvertCurrencies(context.Background(), &cpb.ConvertCurrenciesRequest{From: "BRL", To: "USD", Amount: 25})

	if err != nil {
		log.Fatalf("Error calling server: %s", err)
	}

	log.Printf("Response from call: %s", res)
}
