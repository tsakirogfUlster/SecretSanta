package rest

import (
	"SecretSanta/pkg/models"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) AddMember(ctx *gin.Context) {
	var request models.ExchangeMember
	//This is a good place to use validator for every incoming data. SQL injection but also smelly data should be blocked here.

	// Binding JSON from body to struct
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// call service
	if request.ID == "" || request.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID and Name are required"})
		return
	}

	_, err := c.service.AddMember(request.ID, request.Name)
	if err != nil {
		if errors.Is(err, models.ErrMemberAlreadyExists) {
			ctx.JSON(http.StatusConflict, gin.H{"error": "Member already exists"})
			return
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}
