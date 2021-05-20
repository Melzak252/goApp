package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handlerGET(w http.ResponseWriter, r *http.Request) {

	files, err := ioutil.ReadDir("./files")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Fprintln(w, f.Name())
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
