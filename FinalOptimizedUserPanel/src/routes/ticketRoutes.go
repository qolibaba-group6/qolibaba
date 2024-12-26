
        package routes

        import "net/http"

        
        func RegisterTicketRoutes() {
            http.HandleFunc("/buy-ticket", func(w http.ResponseWriter, r *http.Request) {
                
            })
            http.HandleFunc("/return-ticket", func(w http.ResponseWriter, r *http.Request) {
                
            })
        }
    