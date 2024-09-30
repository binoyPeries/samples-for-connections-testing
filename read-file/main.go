package main

import (
    "io/ioutil"
    "log"
	"os"
	"net/http"
)

func main() {
	http.HandleFunc("/file", fileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	filename := os.Getenv("FILE_PATH")
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(string(content)))
}