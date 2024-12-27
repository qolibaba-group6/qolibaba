
package main

import (
    "encoding/json"
    "net/http"
    "log"
)

func main() {
    mux := http.NewServeMux()

    // Public login endpoint
    mux.HandleFunc("/login", LoginHandler)

    // Protected endpoint
    mux.Handle("/protected", AuthMiddleware(http.HandlerFunc(ProtectedHandler)))

    log.Println("Auth service is running on port 8001...")
    if err := http.ListenAndServe(":8001", mux); err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}

// LoginHandler generates a JWT for a valid user
func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var credentials struct {
        UserID   string `json:"user_id"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Dummy user authentication (replace with actual user validation)
    if credentials.UserID != "user1" || credentials.Password != "password123" {
        http.Error(w, "Invalid user credentials", http.StatusUnauthorized)
        return
    }

    // Generate a token
    token, err := GenerateToken(credentials.UserID)
    if err != nil {
        http.Error(w, "Could not generate token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "token": token,
    })
}

// ProtectedHandler is an example of a protected endpoint
func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
    userID := r.Context().Value("userID").(string)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Welcome to the protected endpoint!",
        "user_id": userID,
    })
}

    // Prometheus metrics endpoint
    mux.Handle("/metrics", MetricsHandler())
    