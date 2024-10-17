package main

import (
	"net/http"
	"user-registration/app"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	myApp := app.NewApp()

	// lambda.Start(myApp.ApiHandler.RegisterUserHandler)

	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.Path {
		case "/register":
			return myApp.ApiHandler.RegisterUserHandler(request)
		case "/login":
			return myApp.ApiHandler.LoginUser(request)
		default:
			return events.APIGatewayProxyResponse{
				Body:       "Not found",
				StatusCode: http.StatusNotFound,
			}, nil
		}
	})

}
