package service

import (
	"file_sharing/fileutil"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func HandleFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleFileGet(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleFileGet(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	inputFilePath := vars["file-name"]

	if inputFilePath == "" {
		http.Error(w, "File name is empty", http.StatusBadRequest)
	}

	file, err := os.Open(filepath.Join(fileutil.Properties["database.path"], inputFilePath))
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Set the appropriate headers
	w.Header().Set("Content-Disposition", "attachment; filename="+inputFilePath)
	w.Header().Set("Content-Type", "application/octet-stream")

	// Copy the file content to the ResponseWriter
	http.ServeContent(w, r, inputFilePath, time.Now(), file)
}
