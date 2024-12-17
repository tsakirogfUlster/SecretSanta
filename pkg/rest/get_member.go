package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *SantaController) GetMember(ctx *gin.Context) {
	ctx.JSON(http.StatusServiceUnavailable, nil)
}
