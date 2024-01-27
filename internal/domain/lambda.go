package domain

// Event es la estructura de entrada para la función Lambda
type Event struct {
	Name string `json:"name"`
}

// Response es la estructura de salida para la función Lambda
type Response struct {
	Message string `json:"message"`
}
