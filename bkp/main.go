package bkp

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/initDb", initDb)
	router.GET("/albums", getAlbums)
	//router.GET("/albums/:id", getAlbumByID)
	//router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
