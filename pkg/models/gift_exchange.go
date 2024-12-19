package models

type GiftExchange struct {
	MemberID          string `json:"member_id"`
	RecipientMemberID string `json:"recipient_member_id"`
}

type GiftExchangeSession struct {
	GiftHistory map[string][]string // Key: member_id, Value: list of recipients in the past years
}
