package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chwnsng/Guessing-Game/backend/models"
	"github.com/chwnsng/Guessing-Game/backend/utils"
)

func GuessHandler(w http.ResponseWriter, r *http.Request) {

	// only allow POST for guessing
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// decode request body into GuessRequest struct
	var req models.GuessRequest
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req) // Decode expects a pointer to the variable we want to populate
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid guess")
		return
	}

	// check if the guess exceeds the pool of possible guesses
	if req.Number < 1 || req.Number > utils.GuessSize {
		utils.RespondError(w, http.StatusOK, fmt.Sprintf("Guesses should be a number between 1 and %v", utils.GuessSize))
		return
	}

	// check guess
	secretNumber := utils.GetSecretNumber()
	if req.Number == secretNumber {
		// returns status 201 on correct guess
		utils.RespondJSON(w, http.StatusCreated, models.GuessResponse{
			Message: "Correct!",
			Correct: true,
		})
		utils.GenerateSecretNumber() // regenerate a new secret number
	} else {
		utils.RespondError(w, http.StatusOK, "Not quite! Try again")
	}

	// close the request body
	defer r.Body.Close()
}
