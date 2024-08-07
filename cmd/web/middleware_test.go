package main

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

//How to test middleware

func Test_addIpToContext(t *testing.T) {

	var app handler

	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false}, // ignore address
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	// Create a new Gin router

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

	for _, e := range tests {
		app.router = gin.Default()
		// Apply the middleware
		app.router.Use(app.addIpToContext())
		// Register the route with the middleware and handler
		app.router.GET("/", nextHandler)

		// Create a new HTTP request
		req := httptest.NewRequest("GET", "http://testing", nil)
		if e.emptyAddr {
			req.RemoteAddr = ""
		}
		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}
		// Record the response
		w := httptest.NewRecorder()

		// Serve the request
		app.router.ServeHTTP(w, req)

	}

}

func Test_ipFromContext(t *testing.T) {

	var app handler
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextUserKey, "whatever")
	expectedIP := ctx.Value(contextUserKey)
	gottenip := app.ipFromContext(ctx)
	t.Log(gottenip)
	if expectedIP != gottenip {
		t.Errorf("expected %v but got %v", expectedIP, gottenip)
	}

}
