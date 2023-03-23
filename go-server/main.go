package main

import (
	"fmt"
	"log"
	"net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Parseform() err:%v", err)
		return
	}
	fmt.Fprintf(w, "Post request successful\n")
	name := r.FormValue("name")
	Email := r.FormValue("e-mail")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", Email)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formhandler)
	http.HandleFunc("/hello", hellohandler)

	fmt.Printf("Starting server at port 8080\n")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
