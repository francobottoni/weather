package repositories

import (
	"context"
	"fmt"
	"github.com/francobottoni/weather/internal/ports"
	"github.com/go-resty/resty/v2"
)

type repository struct {
	ApiKey     string
	ClientRest *resty.Client
}

func NewOpenWeatherClient(apiKey string) ports.OpenWeatherRepo {
	client := resty.New()
	return &repository{ApiKey: apiKey, ClientRest: client}
}

func (c *repository) GetWeatherByCity(ctx context.Context, city string) (*resty.Response, error) {
	apiURL := fmt.Sprintf(urlWeather, city, c.ApiKey)

	resp, err := c.ClientRest.R().Get(apiURL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
