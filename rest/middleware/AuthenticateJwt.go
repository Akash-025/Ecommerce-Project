package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		header := r.Header.Get("Authorization")

		if header == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		headerArr := strings.Split(header, " ")
		if len(headerArr) != 2 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		parsedJwt := strings.Split(accessToken, ".")

		if len(parsedJwt) != 3 {
			http.Error(w, "Authorization Header missing, Unauthorized", http.StatusUnauthorized)
			return
		}
		jwtHeader := parsedJwt[0]
		jwtPayload := parsedJwt[1]
		jwtSignature := parsedJwt[2]

		message := jwtHeader + "." + jwtPayload
		byteArrMsg := []byte(message)
		byteArrSec := []byte(m.cnf.JwtSecretKey)

		h := hmac.New(sha256.New, byteArrSec)
		h.Write(byteArrMsg)
		hash := h.Sum(nil)
		newSignature := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(hash)

		if newSignature != jwtSignature {
			http.Error(w, "Authorization Header missing, Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
