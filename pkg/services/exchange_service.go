package services

import (
	"SecretSanta/pkg/helpers"
	"SecretSanta/pkg/models"
	"SecretSanta/pkg/storage"
	"errors"
	"math/rand"
)

type ExchangeService struct {
	store       *storage.ExchangeStorage
	giftSession *models.GiftExchangeSession
}

func NewExchangeService() *ExchangeService {
	return &ExchangeService{
		store: storage.GetStorage(),
		giftSession: &models.GiftExchangeSession{
			GiftHistory: make(map[string][]string),
		},
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

func (s *ExchangeService) GetGiftExchange() ([]models.GiftExchange, error) {
	// Retrieve members
	members := s.ListMembers()
	if len(members) < 2 {
		return nil, errors.New("at least two members are required for a gift exchange")
	}

	// Prepare data for gift assignment
	memberIDs := make([]string, 0, len(members))
	for id := range members {
		memberIDs = append(memberIDs, id)
	}

	rand.Shuffle(len(memberIDs), func(i, j int) {
		memberIDs[i], memberIDs[j] = memberIDs[j], memberIDs[i]
	})

	// Assign recipients
	assignments := make([]models.GiftExchange, 0, len(memberIDs))
	used := make(map[string]bool) // Tracks assigned recipients
	for i, memberID := range memberIDs {
		recipientID := ""
		for j := 0; j < len(memberIDs); j++ {
			potentialRecipient := memberIDs[(i+j+1)%len(memberIDs)] // Avoid self-gifting
			if !used[potentialRecipient] &&
				potentialRecipient != memberID &&
				!helpers.Contains(s.giftSession.GiftHistory[memberID], potentialRecipient) {
				recipientID = potentialRecipient
				break
			}
		}

		if recipientID == "" {
			return nil, errors.New("unable to assign gift recipients under the current constraints")
		}

		used[recipientID] = true
		assignments = append(assignments, models.GiftExchange{
			MemberID:          memberID,
			RecipientMemberID: recipientID,
		})
		s.giftSession.GiftHistory[memberID] = append(s.giftSession.GiftHistory[memberID], recipientID) // Update history
	}

	return assignments, nil
}
