package main

import (
	"context"
	"log"

	cpb "github.com/rellyson/golang-grpc/proto"
	"github.com/rellyson/golang-grpc/server/services"
	"google.golang.org/grpc"
)

// server is used to implement pb Server.
type server struct {
	cpb.UnimplementedCurrencyServiceServer
}

func SetHandlers(s *grpc.Server) {
	cpb.RegisterCurrencyServiceServer(s, &server{})
}

// Implements GetCurrencies Service from protocol buffer
func (s *server) GetCurrencies(ctx context.Context, in *cpb.CurrenciesRequest) (*cpb.CurrenciesResponse, error) {
	log.Printf("Received: %v", in.String())
	res, err := services.ListCurrencies()

	if err != nil {
		log.Fatal(err)
	}

	return &cpb.CurrenciesResponse{
		Success: res.Success,
		Base:    res.Base,
		Date:    res.Date,
		Rates:   res.Rates,
	}, nil
}
