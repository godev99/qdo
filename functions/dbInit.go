package example

import (
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	message := "Hello world!"

	w.Write([]byte(message))
}
