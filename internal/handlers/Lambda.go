package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/francobottoni/weather/cmd/api"
	"github.com/francobottoni/weather/internal/ports"
	"github.com/francobottoni/weather/pkg/constants"
	"net/http"
)

/*
// HandleLambdaEvent es la función principal que maneja el evento Lambda

	func HandleLambdaEvent(ctx context.Context, event domain.Event) (domain.Response, error) {
		return domain.Response{
			Message: fmt.Sprintf("Hola desde la pc de Franco, %s!", event.Name),
		}, nil
	}
*/

type HandlerLambda struct {
	openWeatherService ports.OpenWeatherService
}

func NewHandlerLambda(openWeatherService ports.OpenWeatherService) *HandlerLambda {
	return &HandlerLambda{
		openWeatherService: openWeatherService,
	}
}

// Este metodo es el por defecto punto de entrada, por lo que en este metodo nosotros realizamos la inyeccion de dependencias y ruteo a los servicios
func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	d := api.IntializeApp()
	ctx := context.Background()
	if request.HTTPMethod == http.MethodGet {
		paramValue, paramExists := request.QueryStringParameters["city"]
		if paramExists {

			weatherResponse, err := d.LambdaHandler.openWeatherService.GetByCity(ctx, paramValue)
			if err != nil {
				return events.APIGatewayProxyResponse{Body: constants.ErrorWhenGetInfo, StatusCode: http.StatusNotFound}, err
			}

			result, _ := json.Marshal(weatherResponse)
			return events.APIGatewayProxyResponse{Body: string(result), StatusCode: http.StatusOK}, nil
		}

		return events.APIGatewayProxyResponse{Body: constants.ResponseDefaultGet, StatusCode: http.StatusOK}, nil
	}

	if request.HTTPMethod == http.MethodPost {
		fmt.Println(request.Body)
		return events.APIGatewayProxyResponse{Body: constants.ResponseDefaultPost, StatusCode: http.StatusOK}, nil
	}

	return events.APIGatewayProxyResponse{Body: constants.NotAllowedMethod, StatusCode: http.StatusBadGateway}, errors.New("execute a allow http.Method")
}
