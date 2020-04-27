package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// FileUploadStatus holds the status of the file upload function
type FileUploadStatus struct {
	Response string `json:"response"`
}

func imgUpload(w http.ResponseWriter, r *http.Request) {

	var data FileUploadStatus

	imgDir, ok := os.LookupEnv("IMG_FOLDER")
	if ok != true {
		log.Fatal("Image directory could not be read.")
	}

	var Buf bytes.Buffer

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	filename := fmt.Sprintf("%s/%s", imgDir, header.Filename)
	// Copy the file data to my buffer
	io.Copy(&Buf, file)

	err = ioutil.WriteFile(filename, Buf.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}

	Buf.Reset()

	imgUploaded.Inc()
	data.Response = "ok"
	json.NewEncoder(w).Encode(data)

}
