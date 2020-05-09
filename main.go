package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "XTRadio CDN.")
	log.Println(r.RemoteAddr, r.Method, r.URL)
}

func cacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
		h.ServeHTTP(w, r)
	})
}

func publishAPI() {
	apiRouter := mux.NewRouter().StrictSlash(true)
	imgDir, ok := os.LookupEnv("IMG_FOLDER")
	if ok != true {
		log.Fatal("Image directory could not be read.")
	}
	apiRouter.HandleFunc("/", homePage)

	apiRouter.PathPrefix("/tracks/").Handler(http.StripPrefix("/tracks/", cacheControlWrapper(http.FileServer(http.Dir(imgDir)))))

	adminServe := mux.NewRouter().StrictSlash(true)
	adminServe.Handle("/metrics", promhttp.Handler())
	adminServe.HandleFunc("/v1/upload", imgUpload).Queries("imgURL", "{imgURL}")

	go func() {
		log.Fatal(http.ListenAndServe(":10001", adminServe))
	}()

	log.Fatal(http.ListenAndServe(":10000", apiRouter))
}

func main() {
	log.Println("XTRadio CDN - v0.1")
	publishAPI()
}
