package rest

import (
	"SecretSanta/pkg/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) AddMember(ctx *gin.Context) {
	var request models.ExchangeMember

	// Binding JSON από το body στο struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Κλήση του service
	if request.ID == "" || request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID and Name are required"})
		return
	}

	c.service.AddMember(request.ID, request.Name)
	ctx.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}
