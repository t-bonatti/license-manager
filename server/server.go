package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t-bonatti/license-manager/controller"
)

// New returns gin.Engine
func New() *gin.Engine {
	c := controller.New()
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.POST("/license", c.Create())
	r.GET("/license/:id/versions/:version", c.Get())

	return r
}
