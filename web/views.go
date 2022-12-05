// Golang
// simple http server
// with html template example
// with json example
// with form example

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type page struct {
	Title    string
	SubTitle string
	Sample1  string
	Sample2  string
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
		return
	}

	page := page{
		Title:    "This is the first page",
		SubTitle: "This is a subtile",
		Sample1:  "another sample var",
		Sample2:  "last sample var",
	}

	w.Header().Set("Content-Type", "text/html")
    //	t, _ := template.ParseFiles("templates/frontpage.html")
	//t.Execute(w, page)
	tpl.ExecuteTemplate(w, "index.html", page)
}

func formHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "POST success")
		name := r.FormValue("firstname")
		address := r.FormValue("lastname")
		fmt.Fprintf(w, "Name = %s\n", name)
		fmt.Fprintf(w, "address = %s\n", address)
		return
	}

	//t, _ := template.ParseFiles("templates/form.html")
	//t.Execute(w, nil)

	tpl.ExecuteTemplate(w, "form.html", nil)
}

func SimpleServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello %s", r.URL.Path[1:])
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello There<h1>")
}

func getBook(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	book := Book{Title: "test title", Author: "Artur Man", Pages: 340}
	json.NewEncoder(w).Encode(book)
}

func sample(w http.ResponseWriter, r *http.Request){
		if r.Method != "POST" {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
		}
}
