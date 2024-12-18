package models

type ExchangeMember struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExchangeMembers map[string]ExchangeMember
