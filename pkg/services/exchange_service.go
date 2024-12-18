package services

import (
	"SecretSanta/pkg/models"
	"SecretSanta/pkg/storage"
	"fmt"
)

type ExchangeService struct {
	store *storage.ExchangeStorage
}

func NewExchangeService() *ExchangeService {
	return &ExchangeService{
		store: storage.GetStorage(), // Shared storage
	}
}

func (s *ExchangeService) AddMember(id, name string) (models.ExchangeMember, error) {
	if id == "" || name == "" {
		return models.ExchangeMember{}, fmt.Errorf("ID and Name cannot be empty")
	}
	member := models.ExchangeMember{ID: id, Name: name}
	s.store.AddMember(member)
	return member, nil
}

func (s *ExchangeService) GetMember(id string) (models.ExchangeMember, error) {
	member, exists := s.store.GetMember(id)
	if !exists {
		return models.ExchangeMember{}, fmt.Errorf("member with ID %s not found", id)
	}
	return member, nil
}

func (s *ExchangeService) DeleteMember(id string) {
	s.store.DeleteMember(id)
}

func (s *ExchangeService) ListMembers() models.ExchangeMembers {
	return s.store.ListMembers()
}
