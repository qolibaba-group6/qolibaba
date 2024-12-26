package main


        import (
            "net/http"
            "strings"

            "github.com/dgrijalva/jwt-go"
        )

        

        package middleware

        func AuthMiddleware(next http.Handler) http.Handler {
            return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                tokenHeader := r.Header.Get("Authorization")
                if tokenHeader == "" {
                    http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
                    return
                }

                tokenParts := strings.Split(tokenHeader, " ")
                if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
                    http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
                    return
                }

                token, err := jwt.Parse(tokenParts[1], func(token *jwt.Token) (interface{}, error) {
                    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                        return nil, http.ErrAbortHandler
                    }
                    return []byte("secret_key"), nil
                })

                if err != nil || !token.Valid {
                    http.Error(w, "Invalid token", http.StatusUnauthorized)
                    return
                }

                next.ServeHTTP(w, r)
            })
        }
    