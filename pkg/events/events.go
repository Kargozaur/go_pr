package events

import "time"

type OrderCreated struct {
	OrderID   string    `json:"orderID"`
	UserID    string    `json:"userID"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type PaymentProceed struct {
	OrderID string `json:"orderID"`
	Status  string `json:"status"` // success | failed | pending
}
