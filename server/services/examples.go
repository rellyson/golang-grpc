package services

import "fmt"

func SayHello(n string) string {
	return fmt.Sprintf("Hello %v!", n)
}

func Sum(num1 int64, num2 int64) int64 {
	return num1 + num2
}
