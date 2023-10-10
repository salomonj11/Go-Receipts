package api

import (
    "testing"

    "github.com/salomonj11/Go-Receipts/models"
)

func TestCalculatePoints(t *testing.T) {
    tests := []struct {
        name     string
        receipt  models.Receipt
        expected int
    }{
        {
            name: "Target Receipt Test",
            receipt: models.Receipt{
                Retailer:      "Target",
                PurchaseDate:  "2022-01-01",
                PurchaseTime:  "13:01",
                Items: []models.Item{
                    {ShortDescription: "Mountain Dew 12PK", Price: 6.49},
                    {ShortDescription: "Emils Cheese Pizza", Price: 12.25},
                    {ShortDescription: "Knorr Creamy Chicken", Price: 1.26},
                    {ShortDescription: "Doritos Nacho Cheese", Price: 3.35},
                    {ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: 12.00},
                },
                Total: 35.35,
            },
            expected: 28,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            got := CalculatePoints(test.receipt)
            if got != test.expected {
                t.Errorf("expected %d points but got %d points", test.expected, got)
            }
        })
    }
}

