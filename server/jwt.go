package server

import (
	"fmt"
	"github.com/ifubar/wow/entities"
	"net"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const JwtHeader = "Jwt"
const jwtKey = "my-secret-key"
const expire = time.Minute * 10

type Claims struct {
	Task entities.Task
	IP   string
	jwt.RegisteredClaims
}

func checkJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		raw := req.Header.Get(JwtHeader)
		claims := Claims{}
		token, err := jwt.ParseWithClaims(raw, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(fmt.Sprintf("invalid token %v", err)))
			return
		}
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if claims.IP != ip {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("ip mismatch"))
			return

		}
		req = req.WithContext(claims.Task.ToCtx(req.Context()))
		next.ServeHTTP(w, req)
	}
}
