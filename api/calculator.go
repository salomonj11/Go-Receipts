package api

import (
    "math"
    "strings"
    "time"

    "github.com/salomonj11/Go-Receipts/models"
)

func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // Rule: Retailer name alphanumeric character count
    points += len(strings.Join(strings.FieldsFunc(receipt.Retailer, func(c rune) bool {
        return !((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'))
    }), ""))

    // Rule: Items (pairs of 2)
    points += (len(receipt.Items) / 2) * 5

    // Rule: Item description character count (multiples of 3)
    // and 0.2 of item price (rounded up)
    for _, item := range receipt.Items {
        if len(item.ShortDescription)%3 == 0 {
            points += int(math.Ceil(item.Price * 0.2))
        }
    }

    // Rule: Odd purchase day
    date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
    if date.Day()%2 != 0 {
        points += 6
    }

    // Rule: Time between 2-4pm
    purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
    if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
        points += 10
    }

    // Rule: Total is a round dollar amount
    if receipt.Total == math.Floor(receipt.Total) {
        points += 50
    }

    // Rule: Total is a multiple of 0.25
    if math.Mod(receipt.Total, 0.25) == 0 {
        points += 25
    }

    return points
}

