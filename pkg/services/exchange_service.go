package services

import (
	"SecretSanta/pkg/models"
	"SecretSanta/pkg/storage"
)

type ExchangeService struct {
	store *storage.ExchangeStorage
}

func NewExchangeService() *ExchangeService {
	return &ExchangeService{
		store: storage.GetStorage(), // Shared storage
	}
}

// In general, I prefer to have a separate file for each service, as I have for rest commands.
// I kept all the services in one page to help with reading of project.
func (s *ExchangeService) AddMember(id, name string) (models.ExchangeMember, error) {
	if id == "" || name == "" {
		return models.ExchangeMember{}, models.ErrEmptyNameORID
	}
	_, exists := s.store.GetMember(id)
	if exists {
		return models.ExchangeMember{}, models.ErrMemberAlreadyExists
	}
	member := models.ExchangeMember{ID: id, Name: name}
	s.store.AddMember(member)
	return member, nil
}

func (s *ExchangeService) GetMember(id string) (models.ExchangeMember, error) {
	member, exists := s.store.GetMember(id)
	if !exists {
		return models.ExchangeMember{}, models.ErrMemberNotFound
	}
	return member, nil
}

func (s *ExchangeService) DeleteMember(id string) error {
	if _, existence := s.store.GetMember(id); existence {
		s.store.DeleteMember(id)
		return nil
	} else {
		return models.ErrMemberNotFound
	}
}

func (s *ExchangeService) ListMembers() models.ExchangeMembers {
	return s.store.ListMembers()
}

func (s *ExchangeService) EditMember(id string, updated models.ExchangeMember) error {
	// Check if the member exists
	if _, exists := s.store.GetMember(id); !exists {
		return models.ErrMemberNotFound
	}

	// Call the store layer to perform the update
	err := s.store.EditMember(id, updated)
	if err != nil {
		return models.ErrFailedToUpdate
	}

	return nil
}
