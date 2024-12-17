package models

type ExchangeMember struct {
	id   string `json:"id"`
	name string `json:"name"`
}

type ExchangeMembers map[string]ExchangeMember
