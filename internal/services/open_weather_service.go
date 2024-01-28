package services

import (
	"context"
	"encoding/json"
	"github.com/francobottoni/weather/internal/domain"
	"github.com/francobottoni/weather/internal/ports"
	"log"
)

type service struct {
	openWeatherClient ports.OpenWeatherRepo
}

func New(repo ports.OpenWeatherRepo) ports.OpenWeatherService {
	return &service{openWeatherClient: repo}
}

func (s *service) GetByCity(c context.Context, city string) (domain.OpenWeatherResponse, error) {
	var openWeatherResponse domain.OpenWeatherResponse
	//TODO implements logs, metrics, etc

	resp, err := s.openWeatherClient.GetWeatherByCity(c, city)
	if err != nil {
		return openWeatherResponse, err
	}

	err = json.Unmarshal(resp.Body(), &openWeatherResponse)
	if err != nil {
		log.Fatalf("Error al decodificar la respuesta JSON: %v", err)
	}

	return openWeatherResponse, nil
}
