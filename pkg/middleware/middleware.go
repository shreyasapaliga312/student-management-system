package middleware

import (
	"net/http"
	"strings"

	"github.com/anaard/simple-student-management/pkg/utils"
)

func CheckAuth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) < 2 {
			utils.WriteErrorResponse(w, http.StatusForbidden, "Token not provided!")
			return
		}
		token := bearerToken[1]

		_, err := utils.VerifyJWTToken(token)
		if err != nil {
			utils.WriteErrorResponse(w, http.StatusForbidden, err.Error())
			return
		}
		next.ServeHTTP(w, r)
	})
}
