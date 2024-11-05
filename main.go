package main

import (
	"fmt"
	"net/http"
	"os"

	"groupie/backend"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Fprintln(os.Stderr, "check args!!!")
		return
	}
	fmt.Println("this is your port : http://localhost:8080/ ")
	http.HandleFunc("/", backend.HandleHome)
	http.HandleFunc("/Artist/{id}", backend.HandlePage)
	http.HandleFunc("/404", backend.ErrorHandler)
	http.HandleFunc("/frontend/css/", backend.CssHandler)
	http.HandleFunc("/frontend/images/", backend.ImageHandler)
	http.ListenAndServe(":8080", nil)
}
