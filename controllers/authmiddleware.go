package controllers

import (
	"github.com/bradenrayhorn/listable-backend/models"
	"github.com/bradenrayhorn/listable-backend/utils"
	"github.com/gorilla/context"
	"net/http"
	"strings"
)

func getToken(header string) string {
	parts := strings.Split(header, " ")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getToken(r.Header.Get("Authorization"))

		user, err := models.FindTokenUser(token)

		if err == nil {
			context.Set(r, AuthUser, user)
			next.ServeHTTP(w, r)
		} else {
			utils.JsonError("invalid api token", w, http.StatusForbidden)
		}
	})
}
