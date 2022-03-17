package example

import (
	"fmt"
	"go.m3o.com/db"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"

	token := "XXX"

	dbService := db.NewDbService(token)
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

	w.Write([]byte(message))
}
