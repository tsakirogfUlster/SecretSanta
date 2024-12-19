package storage

import (
	"SecretSanta/pkg/models"
	"errors"
)

type MockExchangeStorage struct {
	members models.ExchangeMembers
}

func NewMockExchangeStorage() *MockExchangeStorage {
	return &MockExchangeStorage{
		members: make(models.ExchangeMembers),
	}
}

func (m *MockExchangeStorage) AddMember(member models.ExchangeMember) {
	m.members[member.ID] = member
}

func (m *MockExchangeStorage) GetMember(id string) (models.ExchangeMember, bool) {
	member, exists := m.members[id]
	return member, exists
}

func (m *MockExchangeStorage) DeleteMember(id string) {
	delete(m.members, id)
}

func (m *MockExchangeStorage) ListMembers() models.ExchangeMembers {
	copy := make(models.ExchangeMembers)
	for id, member := range m.members {
		copy[id] = member
	}
	return copy
}

func (m *MockExchangeStorage) EditMember(id string, updated models.ExchangeMember) error {
	if _, exists := m.members[id]; !exists {
		return errors.New("member not found")
	}
	delete(m.members, id)
	m.members[updated.ID] = updated
	return nil
}
