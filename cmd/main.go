package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/francobottoni/weather/internal/handlers"
)

func main() {

	// Inicia la función Lambda
	lambda.Start(handlers.HandleRequest)

}
