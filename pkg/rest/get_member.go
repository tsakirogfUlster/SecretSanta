package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) GetMember(ctx *gin.Context) {
	id := ctx.Param("id")
	member, err := c.service.GetMember(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		return
	}
	ctx.JSON(http.StatusOK, member)
}
