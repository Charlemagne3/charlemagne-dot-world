package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", mime.TypeByExtension("html"))
		file, err := ioutil.ReadFile("./static/html/index.html")
		if err != nil {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintf(w, string(file))
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
