package main

import (
	"context"
	"log"

	cpb "github.com/rellyson/golang-grpc/proto/gen-code/currencies"
	e "github.com/rellyson/golang-grpc/proto/gen-code/examples"
	"github.com/rellyson/golang-grpc/server/services"
	"google.golang.org/grpc"
)

// server is used to implement pb Server.
type server struct {
	cpb.UnimplementedCurrencyServiceServer
	e.UnimplementedExampleServiceServer
}

func SetHandlers(s *grpc.Server) {
	cpb.RegisterCurrencyServiceServer(s, &server{})
	e.RegisterExampleServiceServer(s, &server{})
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

// Implements Hello Service from protocol buffer

func (s *server) SayHello(ctx context.Context, in *e.HelloRequest) (*e.HelloResponse, error) {
	log.Printf("Received request to say hello to %v", in.Name)
	m := services.SayHello(in.Name)

	return &e.HelloResponse{
		Message: m,
	}, nil
}

// Implements Sum Service from protocol buffer
func (s *server) Sum(ctx context.Context, in *e.SumRequest) (*e.SumResponse, error) {
	log.Printf("Received request to sum %v and %v", in.Num1, in.Num2)
	r := services.Sum(in.Num1, in.Num2)

	return &e.SumResponse{
		Result: r,
	}, nil
}
