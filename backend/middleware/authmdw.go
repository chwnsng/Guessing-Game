package middleware

import (
	"net/http"
	"strings"

	"github.com/chwnsng/Guessing-Game/backend/utils"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check for Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		// check if the header starts with Bearer
		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid format. Bearer token expected")
			return
		}

		// extract the token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// validate the token
		err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.RespondError(w, http.StatusUnauthorized, "Invalid or expired token")
			return
		}

		// calling the next handler (guessing)
		next.ServeHTTP(w, r)
	}
}
