package domain

type OpenWeatherRequest struct {
	city string `json:"city"`
}

type OpenWeatherResponse struct {
	coord    coord     `json:"coord"`
	timezone int       `json:"timezone"`
	name     string    `json:"name"`
	weather  []weather `json:"weather"`
}

type coord struct {
	lat float32 `json:"lat"`
	lon float32 `json:"lon"`
}

type weather struct {
	description string `json:"description"`
	main        string `json:"main"`
}
