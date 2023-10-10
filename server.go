package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "github.com/salomonj11/Go-Receipts/api"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/receipts/process", api.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", api.GetReceiptPoints).Methods("GET")

    http.Handle("/", r)

    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

