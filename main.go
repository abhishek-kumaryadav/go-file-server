package main

import (
	"file_sharing/fileutil"
	"file_sharing/service"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	fileutil.InitProperties()

	r := mux.NewRouter()
	r.HandleFunc("/file-sharing", service.HandleFileSharing)
	r.HandleFunc("/file-sharing/file/{file-name}", service.HandleFile)
	r.HandleFunc("/file-sharing/search/{file-name}", service.SearchFile)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
