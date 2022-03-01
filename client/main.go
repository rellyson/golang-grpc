package main

import (
	"context"
	"log"
	"os"

	cpb "github.com/rellyson/golang-grpc/proto"
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

	res, err := c.GetCurrencies(context.Background(), &cpb.CurrenciesRequest{})

	if err != nil {
		log.Fatalf("Error calling GetCurrencies Procedure Call: %s", err)
	}

	log.Printf("Response from GetCurrencies Procedure Call: %s", res)
}
