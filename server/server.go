package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t-bonatti/license-manager/controller"
	"github.com/t-bonatti/license-manager/datastore"
)

func New(ds datastore.DataStore) *gin.Engine {

	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.POST("/license", controller.Create(ds))
	r.GET("/license/:id/versions/:version", controller.Get(ds))

	return r
}
