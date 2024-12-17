package models

type GiftExchange struct {
	member_id           string `json:"member_id"`
	recipient_member_id string `json:"recipient_member_id"`
}

type GiftExchangeSession struct {
	gift_exchange map[string]GiftExchange
}
