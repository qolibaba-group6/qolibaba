
package main

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"
    "github.com/gorilla/mux"
)

// CreateUser creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    id, err := CreateUserDB(user.Name, user.Email)
    if err != nil {
        http.Error(w, "Could not create user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User created successfully",
        "user_id": id,
    })
}

// GetUser retrieves a user by ID
func GetUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    user, err := GetUserDB(id)
    if err != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

// UpdateUser updates an existing user's details
func UpdateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    var user struct {
        Name  string `json:"name"`
        Email string `json:"email"`
    }
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = UpdateUserDB(id, user.Name, user.Email)
    if err != nil {
        http.Error(w, "Could not update user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User updated successfully",
    })
}

// DeleteUser removes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    err = DeleteUserDB(id)
    if err != nil {
        http.Error(w, "Could not delete user", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User deleted successfully",
    })
}

func main() {
    InitializeDB() // Initialize the database

    router := mux.NewRouter()

    // CRUD endpoints
    router.Handle("/user", AuthMiddleware(http.HandlerFunc(CreateUser))).Methods("POST")
    router.Handle("/user/{id:[0-9]+}", AuthMiddleware(http.HandlerFunc(GetUser))).Methods("GET")
    router.Handle("/user/{id:[0-9]+}", AuthMiddleware(http.HandlerFunc(UpdateUser))).Methods("PUT")
    router.Handle("/user/{id:[0-9]+}", AuthMiddleware(http.HandlerFunc(DeleteUser))).Methods("DELETE")

    log.Println("User service is running on port 8003...")
    if err := http.ListenAndServe(":8003", router); err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}
