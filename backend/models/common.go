package models

// format for /login POST request body
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// format for /login response (success)
type LoginResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

// default format for error responses
type ErrorResponse struct {
	Message string `json:"message"`
}

type GuessRequest struct {
	Number int `json:"number"`
}

type GuessResponse struct {
	Message string `json:"message"`
	Correct bool   `json:"correct"`
}
