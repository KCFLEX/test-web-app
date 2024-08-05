package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

//How to test middleware

func Test_addIpToContext(t *testing.T) {

	var app handler

	// Create a new Gin router
	app.router = gin.Default()

	// Apply the middleware
	app.router.Use(app.addIpToContext())

	// Define a handler to check the context value
	nextHandler := gin.HandlerFunc(func(ctx *gin.Context) {
		ip := ctx.Request.Context().Value(contextUserKey)
		if ip == nil {
			t.Error("no ip found")
			return
		}

		ipStr, ok := ip.(string)
		if !ok {
			t.Errorf("%v is not a string", ip)
			return
		}

		t.Log(ipStr)
	})

	// Register the route with the middleware and handler
	app.router.GET("/", nextHandler)
	expextedIP := "192.168.1.1"
	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/", nil)

	req.RemoteAddr = "192.168.1.1"

	// Record the response
	w := httptest.NewRecorder()

	// Serve the request
	app.router.ServeHTTP(w, req)

	// Check the response status code
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	if expextedIP != req.RemoteAddr {
		t.Errorf("expected %v, got %q", expextedIP, req.RemoteAddr)
	}
}
