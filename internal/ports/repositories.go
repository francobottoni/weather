package ports

import (
	"context"
	"github.com/go-resty/resty/v2"
)

type OpenWeatherRepo interface {
	GetWeatherByCity(ctx context.Context, city string) (*resty.Response, error)
}
