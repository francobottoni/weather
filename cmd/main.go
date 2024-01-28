package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/francobottoni/weather/cmd/api"
)

func main() {

	// Inicializamos las dependencias
	d := api.InitializeDependencies()

	// Inicia la función Lambda
	// Pasar la función HandleRequest con las dependencias a lambda.Start
	lambda.Start(func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		return d.LambdaHandler.HandleRequest(request)
	})

}
