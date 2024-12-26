package main


        import (
            "net/http"
        )

        

        package routes

        func RegisterAuthRoutes() {
            http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
                
            })

            http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
                
            })
        }
    