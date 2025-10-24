package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"npm-registry/constants"
	"os"
	"strings"
)

func ServeTarBall(w http.ResponseWriter, r *http.Request) {
	log.Println("Authorization header", r.Header.Get("Authorization"))
	log.Println("Incoming tarball request:", r.URL.Path)
	parts := strings.Split(strings.TrimPrefix(r.URL.Path, "/"), "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid tarball URL", http.StatusBadRequest)
		return
	}

	packageName := parts[0]
	tarballFilename := parts[2]

	tarballPath := fmt.Sprintf("/home/aravind-14205/Desktop/npm/npm_test_data/v1/%s/%s", packageName, tarballFilename)
	log.Printf("Serving tarball: %s", tarballPath)
	if _, err := os.Stat(tarballPath); os.IsNotExist(err) {
		log.Printf("Tarball not found: %s", tarballPath)
		http.Error(w, "Tarball not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", tarballFilename))
	http.ServeFile(w, r, tarballPath)
	log.Printf("Successfully served tarball: %s", tarballFilename)
}

func ServeMetaData(w http.ResponseWriter, r *http.Request) {
	log.Println("Authorization header", r.Header.Get("Authorization"))
	log.Println("In coming URI ,", r.URL)
	data, err := os.ReadFile("/home/aravind-14205/Desktop/npm/npm_test_data/v1/hello/meta.json")
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to read meta ", http.StatusNotFound)
		return
	}
	var metadata constants.InstallPackage
	json.Unmarshal(data, &metadata)
	log.Println(metadata)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metadata)
}
