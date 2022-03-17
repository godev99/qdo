package functions

import (
	"fmt"
	//"github.com/gin-gonic/gin"
	"go.m3o.com/db"
	//"net/http"
	"os"
	//"strconv"
)

func InitDb() {
	dbService := db.NewDbService(os.Getenv("M3O_API_TOKEN"))
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
		fmt.Println(err)
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
		fmt.Println(err)
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
		fmt.Println(err)
	}
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
