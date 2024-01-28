package ports

import (
	"context"
	"github.com/francobottoni/weather/internal/domain"
)

type OpenWeatherService interface {
	GetByCity(c context.Context, city string) (domain.OpenWeatherResponse, error)
}
