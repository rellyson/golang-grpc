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
	log.Printf("Received: %v", in.ProtoReflect().Descriptor())
	res, err := services.ListCurrencies()

	if err != nil {
		return &cpb.CurrenciesResponse{}, err
	}

	return &cpb.CurrenciesResponse{
		Success: res.Success,
		Base:    res.Base,
		Date:    res.Date,
		Rates:   res.Rates,
	}, nil
}

// Implements ConnverCurrencies Service from protocol buffer
func (s *server) ConvertCurrencies(ctx context.Context, in *cpb.ConvertCurrenciesRequest) (*cpb.ConvertCurrenciesResponse, error) {
	log.Printf("Received request to convert %v from %v to %v", in.Amount, in.From, in.To)
	res, err := services.ConvertCurrencies(in.From, in.To, in.Amount)

	if err != nil {
		return &cpb.ConvertCurrenciesResponse{}, err
	}

	return &cpb.ConvertCurrenciesResponse{
		Success:    res.Success,
		Info:       res.Info,
		Historical: res.Historical,
		Date:       res.Date,
		Result:     res.Result,
	}, nil
}
