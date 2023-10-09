package main

import (
    "sync"
)

var (
    mu = &sync.Mutex{}

    receiptsMap = make(map[int]models.Receipt)
    nextID = 1
)
