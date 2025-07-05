package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/chwnsng/Guessing-Game/backend/models"
	"github.com/chwnsng/Guessing-Game/backend/utils"
)

// main login logic
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	// Only allows POST for guessing
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// decode request body into LoginRequest struct
	var req models.LoginRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req) // Decode expects a pointer to the variable we want to populate
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Request has invalid payload")
		return
	}

	// check username & pw
	if req.Username == "test" && req.Password == "1234" {
		// generate jwt token
		tokenString, err := utils.CreateToken(req.Username)
		if err != nil {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to generate token")
			return
		}

		utils.RespondJSON(w, http.StatusOK, models.LoginResponse{
			Message: "Log in successful!",
			Token:   tokenString,
		})

	} else {
		utils.RespondError(w, http.StatusUnauthorized, "Username or password incorrect")
	}

	// close the request body
	defer r.Body.Close()
}
