package services

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	cpb "github.com/rellyson/golang-grpc/proto/gen-code/currencies"
)

func ListCurrencies() (*cpb.CurrenciesResponse, error) {
	client := &http.Client{}

	api_key := os.Getenv("CURRENCY_API_KEY")
	url := "http://api.exchangeratesapi.io/v1/latest"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
		return &cpb.CurrenciesResponse{}, err
	}

	q := req.URL.Query()
	q.Add("access_key", api_key)
	req.URL.RawQuery = q.Encode()

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	res, err := client.Do(req)
	if err != nil {
		return &cpb.CurrenciesResponse{}, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer res.Body.Close()

	c := cpb.CurrenciesResponse{}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &c)

	log.Printf("exchange response: %s", body)

	return &c, nil
}

func ConvertCurrencies(f string, t string, a int64) (*cpb.ConvertCurrenciesResponse, error) {
	client := &http.Client{}

	api_key := os.Getenv("CURRENCY_API_KEY")
	url := "http://api.exchangeratesapi.io/v1/convert"

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
		return &cpb.ConvertCurrenciesResponse{}, err
	}

	q := req.URL.Query()
	q.Add("access_key", api_key)
	q.Add("from", f)
	q.Add("to", t)
	q.Add("amount", fmt.Sprint(a))
	req.URL.RawQuery = q.Encode()

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	res, err := client.Do(req)
	if err != nil {
		return &cpb.ConvertCurrenciesResponse{}, err
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer res.Body.Close()

	c := cpb.ConvertCurrenciesResponse{}
	body, _ := io.ReadAll(res.Body)
	json.Unmarshal(body, &c)

	log.Printf("exchange conversion response: %s", body)

	return &c, nil
}
