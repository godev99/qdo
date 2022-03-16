package bkp

import (
	"github.com/gin-gonic/gin"
	"net/http"

	//"github.com/gin-gonic/gin"
	"go.m3o.com/db"
	//"net/http"
	"os"
	//"strconv"
)

// getAlbums responds with the list of all albums as JSON.
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

/*
// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func GetUserbyId(id int) {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
	query := "id == " + strconv.Itoa(id)
	record, err := dbService.Read(&db.ReadRequest{
		Query: query,
		Table: "tdb",
	})

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(record.Records)
}


*/
