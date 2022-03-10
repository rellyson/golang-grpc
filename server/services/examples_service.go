package services

import "fmt"

func SayHello(n string) string {
	return fmt.Sprintf("Hello %v!", n)
}
