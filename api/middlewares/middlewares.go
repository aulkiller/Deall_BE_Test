package middlewares

import (
	"errors"
	"net/http"

	"github.com/aulkiller/Deall_BE_Test/api/auth"
	"github.com/aulkiller/Deall_BE_Test/api/responses"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}

func SetMiddlewareCookieAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// // If use session cookie - Unused ATM replaced by JWT
		// role, err := auth.GetSession(r)
		// if err != nil {
		// 	responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		// 	return
		// }
		userid, role, err := auth.ExtractTokenAttr(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized - Admin access only"))
			return
		}
		if userid != 0 && role != "Admin" {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized - Admin access only"))
			return
		}
		next(w, r)
	}
}
