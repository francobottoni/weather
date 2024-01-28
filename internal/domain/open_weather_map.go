package domain

// REQUEST

type OpenWeatherRequest struct {
	city string `json:"city"`
}

// RESPONSE

// Coordenadas
type Coord struct {
	Lon float64 `json:"lon"`
	Lat float64 `json:"lat"`
}

// Información sobre el clima
type WeatherInfo struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Datos principales
type MainData struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
}

// Viento
type Wind struct {
	Speed float64 `json:"speed"`
	Deg   int     `json:"deg"`
}

// Nubes
type Clouds struct {
	All int `json:"all"`
}

// Sistema
type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}

// Datos de la respuesta
type OpenWeatherResponse struct {
	Coord      Coord         `json:"coord"`
	Weather    []WeatherInfo `json:"weather"`
	Base       string        `json:"base"`
	Main       MainData      `json:"main"`
	Visibility int           `json:"visibility"`
	Wind       Wind          `json:"wind"`
	Clouds     Clouds        `json:"clouds"`
	Dt         int64         `json:"dt"`
	Sys        Sys           `json:"sys"`
	Timezone   int           `json:"timezone"`
	ID         int           `json:"id"`
	Name       string        `json:"name"`
	Cod        int           `json:"cod"`
}
