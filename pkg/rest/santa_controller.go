package rest

import (
	"SecretSanta/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type SantaController struct {
	config    *config.Config
	validator binding.StructValidator
}

// RouterGroups struct contains the various routers groups we add the middleware to
type RouterGroups struct {
	members *gin.RouterGroup
	excange *gin.RouterGroup
}

// NewSnapshotController constructs a controller for snapshot REST API endpoints and registers endpoints' handlers
// in the gin REST engine.
func NewSnapshotController(
	cfg *config.Config,
	router RouterGroups) *SantaController {
	c := &SantaController{
		config: cfg,
	}

	// Snapshot Endpoints
	router.members.GET("/members", c.GetMember)

	return c
}
