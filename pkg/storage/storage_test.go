package storage

import (
	"SecretSanta/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExchangeStorage_CRUD(t *testing.T) {
	// Initialize storage
	storage := GetStorage()

	// Clean up storage for testing
	storage.members = make(models.ExchangeMembers)

	// Test AddMember
	member1 := models.ExchangeMember{ID: "1", Name: "Alice"}
	member2 := models.ExchangeMember{ID: "2", Name: "Bob"}
	storage.AddMember(member1)
	storage.AddMember(member2)

	// Verify members were added
	assert.Equal(t, 2, len(storage.members))
	assert.Equal(t, "Alice", storage.members["1"].Name)
	assert.Equal(t, "Bob", storage.members["2"].Name)

	// Test GetMember
	retrievedMember, exists := storage.GetMember("1")
	assert.True(t, exists)
	assert.Equal(t, "Alice", retrievedMember.Name)

	// Test GetMember for non-existent member
	_, exists = storage.GetMember("3")
	assert.False(t, exists)

	// Test ListMembers
	members := storage.ListMembers()
	assert.Equal(t, 2, len(members))
	assert.Equal(t, "Alice", members["1"].Name)
	assert.Equal(t, "Bob", members["2"].Name)

	// Test EditMember
	updatedMember := models.ExchangeMember{ID: "1", Name: "Alice Updated"}
	err := storage.EditMember("1", updatedMember)
	assert.NoError(t, err)

	// Verify member was updated
	retrievedMember, exists = storage.GetMember("1")
	assert.True(t, exists)
	assert.Equal(t, "Alice Updated", retrievedMember.Name)

	// Test EditMember for non-existent member (this won't throw an error in the current implementation)
	nonExistentMember := models.ExchangeMember{ID: "3", Name: "Non-existent"}
	err = storage.EditMember("3", nonExistentMember)
	assert.NoError(t, err)

	// Test DeleteMember
	storage.DeleteMember("1")
	_, exists = storage.GetMember("1")
	assert.False(t, exists)

	// Verify remaining members
	members = storage.ListMembers()
	assert.Equal(t, 1, len(members))
	assert.Equal(t, "Bob", members["2"].Name)
}

func TestExchangeStorage_Singleton(t *testing.T) {
	// Verify singleton behavior
	storage1 := GetStorage()
	storage2 := GetStorage()
	assert.Equal(t, storage1, storage2, "GetStorage should return the same instance")
}
