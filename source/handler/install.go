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
	log.Println("Authorization header:", r.Header.Get("Authorization"))
	log.Println("Incoming URI:", r.URL.Path)

	pkgName := strings.TrimPrefix(r.URL.Path, "/")
	if pkgName == "" {
		http.Error(w, "Package name missing in URL", http.StatusBadRequest)
		return
	}

	metaFilePath := fmt.Sprintf("/home/aravind-14205/Desktop/npm/npm_test_data/v1/%s/meta.json", pkgName)

	data, err := os.ReadFile(metaFilePath)
	if err != nil {
		log.Printf("Error reading metadata for package %s: %v", pkgName, err)
		http.Error(w, fmt.Sprintf("Package '%s' not found", pkgName), http.StatusNotFound)
		return
	}

	var metadata constants.Package
	if err := json.Unmarshal(data, &metadata); err != nil {
		log.Printf("Error parsing JSON for package %s: %v", pkgName, err)
		http.Error(w, "Invalid metadata format", http.StatusInternalServerError)
		return
	}

	latest := metadata.DistTags["latest"]
	log.Printf("Serving metadata for package: %s@%s", metadata.Name, latest)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(metadata); err != nil {
		log.Println("Error encoding JSON response:", err)
	}
}
