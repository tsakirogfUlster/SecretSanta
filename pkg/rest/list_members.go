package rest

import (
	"SecretSanta/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) ListMembers(ctx *gin.Context) {
	// Call the service to retrieve the list of members
	membersMap := c.service.ListMembers()

	// Convert the map to a slice for clean JSON output
	membersSlice := make([]models.ExchangeMember, 0, len(membersMap))
	for _, member := range membersMap {
		membersSlice = append(membersSlice, member)
	}

	// Return the list of members as JSON
	ctx.JSON(http.StatusOK, membersSlice)
}
