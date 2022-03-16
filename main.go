package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	//"github.com/gin-gonic/gin"
	"go.m3o.com/db"
	//"net/http"
	"os"
	//"strconv"
)

func main() {
	router := gin.Default()
	router.GET("/initDb", initDb)
	router.GET("/albums", getAlbums)
	//router.GET("/albums/:id", getAlbumByID)
	//router.POST("/albums", postAlbums)

	router.Run("qdo.m3o.app:8080")
}

func initDb(c *gin.Context) {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	dbService.DropTable(&db.DropTableRequest{
		Table: "albums",
	})
	//if err != nil {
	//	c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	//}

	time.Sleep(3)

	_, err := dbService.Create(&db.CreateRequest{
		Record: map[string]interface{}{
			"Id":      "1",
			"Title":   "Blue Train",
			"Artiste": "John Coltrane",
			"Price":   "56.99",
		},
		Table: "albums",
	})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

	_, err = dbService.Create(&db.CreateRequest{
		Record: map[string]interface{}{
			"Id":      "2",
			"Title":   "Jeru",
			"Artiste": "Gerry Mulligan",
			"Price":   "17.99",
		},
		Table: "albums",
	})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}

	_, err = dbService.Create(&db.CreateRequest{
		Record: map[string]interface{}{
			"Id":      "3",
			"Title":   "Sarah Vaughan and Clifford Brown",
			"Artiste": "Sarah Vaughan",
			"Price":   "39.99",
		},
		Table: "albums",
	})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
}

func getAlbums(c *gin.Context) {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	albums, err := dbService.Read(&db.ReadRequest{
		Table: "albums",
	})
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
	c.IndentedJSON(http.StatusOK, albums.Records)
}
