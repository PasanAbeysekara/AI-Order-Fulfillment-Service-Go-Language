package main

type Delivery struct {
	ID      int    `json:"id"`
	OrderID int    `json:"order_id"`
	Status  string `json:"status"`
}
