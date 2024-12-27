
package main

import (
    "encoding/json"
    "net/http"
    "log"
    "strconv"
    "github.com/gorilla/mux"
)

// DepositHandler handles deposit transactions
func DepositHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    walletID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
        return
    }

    var body struct {
        Amount float64 `json:"amount"`
    }
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Amount <= 0 {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = DepositDB(walletID, body.Amount)
    if err != nil {
        http.Error(w, "Could not process deposit", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Deposit successful",
    })
}

// WithdrawHandler handles withdrawal transactions
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    walletID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
        return
    }

    var body struct {
        Amount float64 `json:"amount"`
    }
    if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.Amount <= 0 {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    err = WithdrawDB(walletID, body.Amount)
    if err != nil {
        http.Error(w, "Could not process withdrawal", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "Withdrawal successful",
    })
}

// TransactionHistoryHandler retrieves transaction history
func TransactionHistoryHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    walletID, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
        return
    }

    history, err := GetTransactionHistoryDB(walletID)
    if err != nil {
        http.Error(w, "Could not retrieve transaction history", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(history)
}

func main() {
    InitializeDB() // Initialize the database

    router := mux.NewRouter()

    // Wallet endpoints
    router.Handle("/wallet/{id:[0-9]+}/deposit", AuthMiddleware(http.HandlerFunc(DepositHandler))).Methods("POST")
    router.Handle("/wallet/{id:[0-9]+}/withdraw", AuthMiddleware(http.HandlerFunc(WithdrawHandler))).Methods("POST")
    router.Handle("/wallet/{id:[0-9]+}/transactions", AuthMiddleware(http.HandlerFunc(TransactionHistoryHandler))).Methods("GET")

    log.Println("Wallet service is running on port 8002...")
    if err := http.ListenAndServe(":8002", router); err != nil {
        log.Fatalf("Could not start server: %s", err)
    }
}
