package service

import (
	"encoding/json"
	"file_sharing/fileutil"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Response struct {
	Files []string `json:"files"`
}

func SearchFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inputFilePath := vars["file-name"]

	matches, err := fileutil.SearchFiles(fileutil.Properties["database.path"], inputFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error searching files: %v", err), http.StatusInternalServerError)
		return
	}
	response := Response{
		Files: matches,
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error encoding JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
