package handlers

import (
	"context"
	"fmt"
	"github.com/francobottoni/weather/internal/domain"
)

// HandleLambdaEvent es la función principal que maneja el evento Lambda
func HandleLambdaEvent(ctx context.Context, event domain.Event) (domain.Response, error) {
	return domain.Response{
		Message: fmt.Sprintf("Hola desde la pc de Franco, %s!", event.Name),
	}, nil
}
