package rest

import (
	"SecretSanta/pkg/config"
	"SecretSanta/pkg/services"
	"crypto/tls"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

type REST struct {
	config *config.Config
	errgrp *errgroup.Group

	Engine *gin.Engine
	Routes *gin.RouterGroup
}

func NewREST(c *config.Config,
	service *services.ExchangeService) *REST {
	engine := gin.New()
	routes := engine.Group("")

	memberGroup := routes.Group("/v1")
	exchangeGroup := routes.Group("/v1")

	// TO-DO : Add logger, prometheous metrics to middleware and use them here.
	routes.Use()

	routers := RouterGroups{
		members: memberGroup,
		excange: exchangeGroup,
	}
	NewSantaController(service, c, routers)
	engine.Use(
		gin.Recovery(),
	)

	return &REST{
		Engine: engine,
		Routes: routes,
		config: c,
	}
}

// TLSconfig: pem files are generated and placed in project folder for this proof of concept.
// We need to generate new and place them in a safe environment (i.e. a vault/secrets)
func (r *REST) Run() error {
	server := http.Server{
		Addr:              net.JoinHostPort(r.config.Host, r.config.Port),
		Handler:           r.Engine,
		ReadHeaderTimeout: 60 * time.Second,
		TLSConfig: &tls.Config{
			ServerName: r.config.Host,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_AES_128_GCM_SHA256,
				tls.TLS_AES_256_GCM_SHA384,
				tls.TLS_CHACHA20_POLY1305_SHA256,
			},
			MinVersion: tls.VersionTLS12,
		},
	}

	return server.ListenAndServeTLS("cert.pem", "key.pem")
}
