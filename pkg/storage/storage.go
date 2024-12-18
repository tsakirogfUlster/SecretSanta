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

// GetStorage initializes and returns the singleton storage instance to be shared across the services
// At some point when this project is not an MVP, we'll need to update this with a client that will be responsible to
// use a persistent data store (like postGres, or even GraphDB).
// In that case we still need a singleton for the client of postGress, therefore we have very little harsh to do the update.
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
	// We could here Check if the member exists. It's good practice to have all the errors thrown in service or storage layer, depending where they are created.
	//This is to demostrate that we can have the errors in both layers In this example.
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

func (s *ExchangeStorage) EditMember(id string, updated models.ExchangeMember) error {
	s.Lock()
	defer s.Unlock()
	// Update the member data
	s.members[id] = updated
	return nil
}
