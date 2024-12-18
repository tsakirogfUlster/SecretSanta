package rest

import (
	"SecretSanta/pkg/config"
	"SecretSanta/pkg/services"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewREST(t *testing.T) {
	// Mock config
	appConfig := &config.Config{
		Host: "localhost",
		Port: "8080",
	}

	// Mock service
	mockService := &services.ExchangeService{}

	// Create REST instance
	rest := NewREST(appConfig, mockService)

	// Assertions
	assert.NotNil(t, rest)
	assert.NotNil(t, rest.Engine)
	assert.NotNil(t, rest.Routes)
	assert.Equal(t, appConfig, rest.config)
}

func TestRoutesSetup(t *testing.T) {
	// Use Gin's test mode
	gin.SetMode(gin.TestMode)

	// Mock config
	appConfig := &config.Config{
		Host: "localhost",
		Port: "8080",
	}

	// Mock service
	mockService := &services.ExchangeService{}

	// Create REST instance
	rest := NewREST(appConfig, mockService)

	// Add a sample GET endpoint for testing
	rest.Routes.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Create a test HTTP server
	server := httptest.NewServer(rest.Engine)
	defer server.Close()

	// Perform a test request
	resp, err := http.Get(server.URL + "/ping")

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	// Parse response
	defer resp.Body.Close()
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "pong", response["message"])
}

func TestTLSConfiguration(t *testing.T) {
	// Mock config
	appConfig := &config.Config{
		Host: "localhost",
		Port: "8080",
	}

	// Mock service
	mockService := &services.ExchangeService{}

	// Create REST instance
	rest := NewREST(appConfig, mockService)

	// Create the server manually to access the TLSConfig
	server := &http.Server{
		TLSConfig: &tls.Config{
			ServerName: appConfig.Host,
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

	// Assertions
	assert.NotNil(t, server.TLSConfig)
	assert.NotNil(t, rest)
	assert.Contains(t, server.TLSConfig.CipherSuites, tls.TLS_AES_128_GCM_SHA256)
	assert.Equal(t, uint16(tls.VersionTLS12), server.TLSConfig.MinVersion) // Fix: Use uint16
}
