package api

import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"

    "github.com/gorilla/mux"
    "github.com/salomonj11/Go-Receipts/models"
)

var (
    mu          sync.Mutex
    nextID      int
    receiptsMap = make(map[int]models.Receipt)
)

func ProcessReceipt (w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    } 



    mu.Lock()
    currentID := nextID
    receiptsMap[currentID] = receipt
    nextID++
    mu.Unlock()

    response := models.ReceiptResponse{ID: strconv.Itoa(currentID)}
    json.NewEncoder(w).Encode(response)
}

func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    idStr, ok := vars["id"]
    if !ok {
        http.Error(w, "ID not provided", http.StatusBadRequest)
        return
    }

    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid ID format", http.StatusBadRequest)
        return
    }

    receipt, ok := receiptsMap[id]
    if !ok {
        http.Error(w, "Receipt not found", http.StatusNotFound)
        return
    }

    // Placeholder response for now.
    json.NewEncoder(w).Encode(models.PointsResponse{Points: 0})
}
