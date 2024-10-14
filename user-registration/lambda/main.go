package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Username string `json:"username"`
}

func EventHandler(myEvent MyEvent) (string, error) {
	if myEvent.Username == "" {
		return "", fmt.Errorf("username cannot be empty")
	}

	return fmt.Sprintf("Successful %s", myEvent.Username), nil
}

func main() {
	lambda.Start(EventHandler)
}
