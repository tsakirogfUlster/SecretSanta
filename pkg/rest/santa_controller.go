package rest

import (
	"SecretSanta/pkg/config"
	"SecretSanta/pkg/services"
	"github.com/gin-gonic/gin"
)

type SantaController struct {
	config  *config.Config
	service services.ExchangeService
}

// RouterGroups struct contains the various routers groups we add the middleware to
type RouterGroups struct {
	members *gin.RouterGroup
	excange *gin.RouterGroup
}

// NewSantaController constructs a controller for snapshot REST API endpoints and registers endpoints' handlers
// in the gin REST engine.
func NewSantaController(
	service *services.ExchangeService,
	cfg *config.Config,
	router RouterGroups) *SantaController {
	c := &SantaController{
		config:  cfg,
		service: *service,
	}

	// Snapshot Endpoints
	router.members.GET("/members/:id", c.GetMember)
	router.members.GET("/members/", c.ListMembers)
	router.members.POST("/members", c.AddMember)
	router.members.PUT("/members/:id", c.EditMember)
	router.members.DELETE("/members/:id", c.DeleteMember)
	//router.excange.GET("/gift_exchange/", c.GetGiftExchange)

	return c
}
