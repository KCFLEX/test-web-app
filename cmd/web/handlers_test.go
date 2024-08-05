package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/stretchr/testify/assert"
)

func Test_home(t *testing.T) {

	tests := []struct {
		name               string
		url                string
		expectedStatusCode int
	}{
		{"success", "/", http.StatusOK},
		{"not found", "/dog", http.StatusNotFound},
	}

	app := &handler{}
	app.router = gin.Default()
	app.router.GET("/", app.home)
	pathToTemplates = "./../../templates/"

	for _, e := range tests {
		req, _ := http.NewRequest("GET", e.url, nil)
		w := httptest.NewRecorder()
		// serving the handler
		app.router.ServeHTTP(w, req)
		assert.Equal(t, e.expectedStatusCode, w.Code, "for %s: expected status %d, but got %d", e.name, e.expectedStatusCode, w.Code)
	}

}
