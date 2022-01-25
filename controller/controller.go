package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/t-bonatti/license-manager/datastore"
	"github.com/t-bonatti/license-manager/model"
)

// Create a lincense
func Create(ds datastore.DataStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		var license model.License
		if err := c.Bind(&license); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		license.CreatedAt = time.Now()
		if err := ds.Create(license); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		c.Writer.WriteHeader(http.StatusCreated)
	}
}

// Get a lincense by version
func Get(ds datastore.DataStore) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		version := c.Param("version")

		license, err := ds.Get(id, version)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				c.String(http.StatusNotFound, "Not found")
			} else {
				c.String(http.StatusInternalServerError, err.Error())
			}
			return
		}

		c.JSON(http.StatusOK, license)
	}
}
