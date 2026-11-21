package qrmenu

import "time"

type QRItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	IsSoldOut bool   `json:"is_sold_out"`
}

type KitchenTicket struct {
	TicketID  int       `json:"ticket_id"`
	TableCode string    `json:"table_code"`
	ItemNames []string  `json:"item_names"`
	OrderedAt time.Time `json:"ordered_at"`
	Status    string    `json:"status"` // COOKING, READY
}
