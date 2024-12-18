package storage

import (
	"SecretSanta/pkg/models"
	"sync"
)

// ExchangeStorage defines thread-safe storage for ExchangeMembers
type ExchangeStorage struct {
	sync.RWMutex
	members models.ExchangeMembers
}

var storageInstance *ExchangeStorage
var once sync.Once

// GetStorage initializes and returns the singleton storage instance
func GetStorage() *ExchangeStorage {
	once.Do(func() {
		storageInstance = &ExchangeStorage{
			members: make(models.ExchangeMembers),
		}
	})
	return storageInstance
}

// CRUD Operations
func (s *ExchangeStorage) AddMember(member models.ExchangeMember) {
	s.Lock()
	defer s.Unlock()
	s.members[member.ID] = member
}

func (s *ExchangeStorage) GetMember(id string) (models.ExchangeMember, bool) {
	s.RLock()
	defer s.RUnlock()
	member, exists := s.members[id]
	return member, exists
}

func (s *ExchangeStorage) DeleteMember(id string) {
	s.Lock()
	defer s.Unlock()
	delete(s.members, id)
}

func (s *ExchangeStorage) ListMembers() models.ExchangeMembers {
	s.RLock()
	defer s.RUnlock()
	copy := make(models.ExchangeMembers)
	for id, member := range s.members {
		copy[id] = member
	}
	return copy
}
