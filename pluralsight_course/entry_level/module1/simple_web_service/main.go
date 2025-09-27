package simplewebservice

import (
	"fmt"
	"net/http"
)

func SimpleWebService() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Web Service are easy with GO!")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./simple_web_service/home.html")
	})

	http.ListenAndServe(":3000", nil)
}
