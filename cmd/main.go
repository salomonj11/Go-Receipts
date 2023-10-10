package main

import (
    "sync"
    "github.com/salomonj11/Go-Receipts/models"
)

var (
    mu = &sync.Mutex{}

    receiptsMap = make(map[int]models.Receipt)
    nextID = 1
)
