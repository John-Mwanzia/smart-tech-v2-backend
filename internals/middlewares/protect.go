package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/John-Mwanzia/smart-tech-v2-backend/internals/config"
	"github.com/golang-jwt/jwt"
)

type contextKey string 

const userIdKey = contextKey("id")

func JwtAuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
   
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		//trim header
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token , err := jwt.Parse(tokenStr, func(t *jwt.Token) (any, error) {
      return []byte(config.AppEnv.Jwtsecret), nil 
		})


		if err != nil || !token.Valid {
			http.Error(w, "Invalid token"+ ": " +err.Error(), http.StatusUnauthorized)
			return  
		}


		claims, ok := token.Claims.(jwt.MapClaims)
		
		if !ok || claims["id"] == nil {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return 
		}

		//add context 

		ctx := context.WithValue(r.Context(), userIdKey, claims["id"].(string))

		next.ServeHTTP(w, r.WithContext(ctx))

	}) 

}

//To avoid accessing userIDKey directly, expose a helper
func UserIDKey() contextKey {
	return userIdKey
}


