// Package for request handlers
package controllers

import (
	albumsDb "github.com/danilevy1212/album-service-gin/db"
	"github.com/danilevy1212/album-service-gin/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

var dbClient = albumsDb.New()

// GetAll, return all albums in DDBB
func GetAll(c *gin.Context) {
	albums, error := dbClient.GetAll()

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"albums": albums,
	})
}

// Create a new album
func Create(c *gin.Context) {
	var newAlbum models.AlbumPostBody
	if err := c.ShouldBindJSON(&newAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, err := dbClient.Insert(&newAlbum)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &album)
}

// Get by ID
func GetByID(c *gin.Context) {
	album, notFound := dbClient.GetById(c.Param("id"))

	if notFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": notFound.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

// Delete by ID
func Delete(c *gin.Context) {
	album, notFound := dbClient.Delete(c.Param("id"))

	if notFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": notFound.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}

// Patch by ID
func Patch(c *gin.Context) {
	var patchAlbum models.AlbumPatchBody
	if err := c.ShouldBindJSON(&patchAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album, notFound := dbClient.Patch(c.Param("id"), &patchAlbum)
	if notFound != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": notFound.Error()})
		return
	}

	c.JSON(http.StatusOK, album)
}
