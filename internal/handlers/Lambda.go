package handlers

import (
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
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

func HandleRequest(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.HTTPMethod == http.MethodGet {
		return events.APIGatewayProxyResponse{Body: constants.ResponseDefaultGet, StatusCode: http.StatusOK}, nil
	}

	if request.HTTPMethod == http.MethodPost {
		fmt.Println(request.Body)
		return events.APIGatewayProxyResponse{Body: constants.ResponseDefaultPost, StatusCode: http.StatusOK}, nil
	}

	return events.APIGatewayProxyResponse{Body: constants.NotAllowedMethod, StatusCode: http.StatusBadGateway}, errors.New("execute a allow http.Method")
}
