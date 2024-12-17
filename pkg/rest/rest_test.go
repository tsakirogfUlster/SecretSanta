package rest

import (
	"SecretSanta/pkg/config"
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

	// Create REST instance
	rest := NewREST(appConfig)

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

	// Create REST instance
	rest := NewREST(appConfig)

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
}

func TestRunMethod_NoTLSFiles(t *testing.T) {
	// Mock config
	appConfig := &config.Config{
		Host: "localhost",
		Port: "8080",
	}

	// Create REST instance
	rest := NewREST(appConfig)

	// Mock HTTP server (skip ListenAndServeTLS call)
	go func() {
		err := rest.Run()
		assert.Error(t, err) // Expect an error because cert files are missing
	}()
}
