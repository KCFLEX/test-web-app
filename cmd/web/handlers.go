package main

import (
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

var pathToTemplates = "./templates/"

type PageData struct {
	IP   string
	Data map[string]any
}

func (h *handler) home(ctx *gin.Context) {
	err := h.render(ctx, "home.page.gohtml", &PageData{})
	if err != nil {
		log.Printf("Error rendering template: %v", err)
	}
}

//how to render template from disk in go

func (h *handler) render(ctx *gin.Context, filename string, data *PageData) error {
	temp, err := template.ParseFiles(path.Join(pathToTemplates, filename))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"failed to parse files": err.Error()})
	}
	data.IP = h.ipFromContext(ctx.Request.Context())

	err = temp.Execute(ctx.Writer, data)
	if err != nil {
		return err
	}

	return nil
}
