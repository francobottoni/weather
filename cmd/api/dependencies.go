package api

import (
	"github.com/francobottoni/weather/internal/handlers"
	"github.com/francobottoni/weather/internal/ports"
	"github.com/francobottoni/weather/internal/repositories"
	"github.com/francobottoni/weather/internal/services"
)

const apiKey = "be017b21b3c8c066012a20d27ad6b55c"

type DependenciesProyect struct {
	OpenWeatherRepo    ports.OpenWeatherRepo
	OpenWeatherService ports.OpenWeatherService
	LambdaHandler      *handlers.HandlerLambda
}

func initializeDependencies() *DependenciesProyect {
	dependencies := DependenciesProyect{}

	dependencies.OpenWeatherRepo = repositories.NewOpenWeatherClient(apiKey)
	dependencies.OpenWeatherService = services.New(dependencies.OpenWeatherRepo)
	dependencies.LambdaHandler = handlers.NewHandlerLambda(dependencies.OpenWeatherService)

	return &dependencies
}
