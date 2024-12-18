package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) DeleteMember(ctx *gin.Context) {
	id := ctx.Param("id")

	// Call the service layer to delete the member
	err := c.service.DeleteMember(id)
	if err != nil {
		// Handle the case where the member does not exist
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}

	// Return a success response
	ctx.JSON(http.StatusOK, gin.H{"message": "Member deleted successfully"})
}
