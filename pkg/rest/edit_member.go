package rest

import (
	"SecretSanta/pkg/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) EditMember(ctx *gin.Context) {
	id := ctx.Param("id")

	// Parse the request body into a Member object
	var updatedMember models.ExchangeMember
	if err := ctx.ShouldBindJSON(&updatedMember); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service layer to update the member
	err := c.service.EditMember(id, updatedMember)
	if err != nil {
		if errors.Is(err, models.ErrMemberNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
		}
		return
	}

	// Return success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Member updated successfully"})
}
