package middleware

import (
	"net/http"
	"os"

	"github.com/tshubham7/eth-parser/internal/pkg/constants"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		adminToken := r.Header.Get(constants.HeaderAdminAuthToken)
		if adminToken != os.Getenv(constants.EnvAdminAuthToken) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
