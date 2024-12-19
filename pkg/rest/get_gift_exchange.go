package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) GetGiftExchange(ctx *gin.Context) {
	assignments, err := c.service.GetGiftExchange()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, assignments)
}
