package main

import "github.com/gin-gonic/gin"

type handler struct {
	router *gin.Engine
}

func handlerNew() *handler {
	router := gin.Default()
	return &handler{
		router: router,
	}
}

func (h *handler) registerRoutes() {

	//registering middleware
	h.router.Use(gin.Recovery())
	h.router.Use(h.addIpToContext())

	h.router.GET("/", h.home)
	h.router.Static("/static", "./static")
	h.router.POST("/login", h.login)

}

func (h *handler) serve() error {
	h.registerRoutes()
	return h.router.Run(":8080")
}
