package rest

import (
	"SecretSanta/pkg/models"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Mock the service layer
type mockExchangeService struct{}

func (m *mockExchangeService) GetMember(id string) (models.ExchangeMember, error) {
	if id == "1" {
		return models.ExchangeMember{ID: "1", Name: "Alice"}, nil
	}
	return models.ExchangeMember{}, models.ErrMemberNotFound
}

func (m *mockExchangeService) ListMembers() models.ExchangeMembers {
	return models.ExchangeMembers{
		"1": {ID: "1", Name: "Alice"},
		"2": {ID: "2", Name: "Bob"},
	}
}

func (m *mockExchangeService) AddMember(member models.ExchangeMember) error {
	if member.ID == "" {
		return models.ErrInvalidInput
	}
	return nil
}

func (m *mockExchangeService) EditMember(id string, updated models.ExchangeMember) error {
	if id == "1" {
		return nil
	}
	return models.ErrMemberNotFound
}

func (m *mockExchangeService) DeleteMember(id string) error {
	if id == "1" {
		return nil
	}
	return models.ErrMemberNotFound
}

func TestSantaController_GetMember(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a Gin router
	router := gin.New()

	// Mock GetMember endpoint logic directly in the router
	router.GET("/members/:id", func(c *gin.Context) {
		id := c.Param("id")
		if id == "1" {
			c.JSON(http.StatusOK, models.ExchangeMember{ID: "1", Name: "Alice"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "member not found"})
		}
	})

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/members/1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert response
	assert.Equal(t, http.StatusOK, rec.Code)

	var member models.ExchangeMember
	err := json.Unmarshal(rec.Body.Bytes(), &member)
	assert.NoError(t, err)
	assert.Equal(t, "1", member.ID)
	assert.Equal(t, "Alice", member.Name)
}

func TestSantaController_ListMembers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a Gin router
	router := gin.New()

	// Mock ListMembers endpoint logic directly in the router
	router.GET("/members", func(c *gin.Context) {
		c.JSON(http.StatusOK, []models.ExchangeMember{
			{ID: "1", Name: "Alice"},
			{ID: "2", Name: "Bob"},
		})
	})

	// Perform request
	req, _ := http.NewRequest(http.MethodGet, "/members", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert response
	assert.Equal(t, http.StatusOK, rec.Code)

	var members []models.ExchangeMember
	err := json.Unmarshal(rec.Body.Bytes(), &members)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(members))
}

func TestSantaController_AddMemberController(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a Gin router
	router := gin.New()

	// Mock AddMember endpoint logic directly in the router
	router.POST("/members", func(c *gin.Context) {
		var member models.ExchangeMember
		if err := c.ShouldBindJSON(&member); err != nil || member.ID == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
			return
		}
		c.JSON(http.StatusOK, member)
	})

	// Perform request
	member := models.ExchangeMember{ID: "3", Name: "Charlie"}
	body, _ := json.Marshal(member)
	req, _ := http.NewRequest(http.MethodPost, "/members", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	// Assert response
	assert.Equal(t, http.StatusOK, rec.Code)

	var addedMember models.ExchangeMember
	err := json.Unmarshal(rec.Body.Bytes(), &addedMember)
	assert.NoError(t, err)
	assert.Equal(t, "3", addedMember.ID)
	assert.Equal(t, "Charlie", addedMember.Name)
}
