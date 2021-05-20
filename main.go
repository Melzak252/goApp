package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlerGET(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	html, err := ioutil.ReadFile("home.html")
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(html)
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Fprintln(w, f.Name()+"<br>")
	}
}

func handlerPOST(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(handler.Filename, fileBytes, 0644)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		handlerPOST(w, r)
		handlerGET(w, r)
	} else if r.Method == "GET" {
		handlerGET(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
