package main

import (
		"net/http"
		"log"
)

func main() {

		//fileServer := http.FileServer(http.Dir("./static"))
		mux := http.NewServeMux()
		mux.HandleFunc("/", indexHandler)
		mux.HandleFunc("/book", getBook)
		mux.HandleFunc("/form", formHandler)

		mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

		log.Println("Starting server at port 8000")
		if err := http.ListenAndServe(":8000", mux); err != nil {
				log.Fatal(err)
		}
}
