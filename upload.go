package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// FileUploadStatus holds the status of the file upload function
type FileUploadStatus struct {
	Response string `json:"response"`
	Filename string `json:"filename"`
}

func imgUpload(w http.ResponseWriter, r *http.Request) {

	var data FileUploadStatus
	var filepath string

	data.Response = "ok"

	imgDir, ok := os.LookupEnv("IMG_FOLDER")
	if ok != true {
		log.Fatal("Image directory could not be read - Please set ENV Variables.")
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}

	fileURL := mux.Vars(r)["imgURL"]

	data.Filename, filepath = generateFilename(imgDir)

	err := downloadFile(filepath, fileURL)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	imgUploaded.Inc()
	json.NewEncoder(w).Encode(data)

}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
