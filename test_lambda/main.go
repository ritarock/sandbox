package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func hello() (string, error) {
	fmt.Println("hello")
	return "hello", nil
}

func main() {
	lambda.Start(hello)
}
