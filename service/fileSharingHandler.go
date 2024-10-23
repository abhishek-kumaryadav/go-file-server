package service

import (
	"file_sharing/fileutil"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func HandleFileSharing(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		handleFileSharingGet(w, r)
	case http.MethodPost:
		handleFileSharingPost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleFileSharingGet(w http.ResponseWriter, r *http.Request) {
	filePaths, err := getFolderStructure(fileutil.Properties["database.path"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, filePaths)
}

func handleFileSharingPost(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving the file: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// create a destination file
	dst, _ := os.Create(filepath.Join(fileutil.Properties["database.path"], header.Filename))
	defer dst.Close()

	// upload the file to destination path
	_, _ = io.Copy(dst, file)

	fmt.Println("File uploaded successfully")
}

func getFolderStructure(root string) (string, error) {
	var result strings.Builder
	result.WriteString(filepath.Base(root) + "\n")

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == root {
			return nil // Skip the root directory itself
		}

		relPath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		depth := strings.Count(relPath, string(os.PathSeparator))
		prefix := strings.Repeat("│   ", depth)

		if info.IsDir() {
			result.WriteString(fmt.Sprintf("%s├── %s/\n", prefix, info.Name()))
		} else {
			result.WriteString(fmt.Sprintf("%s├── %s\n", prefix, info.Name()))
		}

		return nil
	})

	return result.String(), err
}
