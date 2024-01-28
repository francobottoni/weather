package domain

type OpenWeatherRequest struct {
	city string `json:"city"`
}

type OpenWeatherResponse struct {
	lat     string      `json:"lat"`
	lon     string      `json:"lon"`
	alerts  interface{} `json:"alerts"`
	current current
}

type current struct {
	weather    weather
	clouds     string `json:"clouds"`
	visibility string `json:"visibility"`
}

type weather struct {
	description string `json:"description"`
	main        string `json:"main"`
	id          string `json:"id"`
}
