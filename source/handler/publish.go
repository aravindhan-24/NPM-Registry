package handler

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"npm-registry/constants"
	"os"
	"path/filepath"
	"strings"
)

func HandlePublish(w http.ResponseWriter, r *http.Request) {
	log.Println("Authorization header:", r.Header.Get("Authorization"))

	httpBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read body:", err)
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}

	var publishStruct constants.Package
	if err := json.Unmarshal(httpBody, &publishStruct); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	packageName := publishStruct.Name
	if packageName == "" {
		http.Error(w, "Package name missing in JSON", http.StatusBadRequest)
		return
	}

	baseDir := "/home/aravind-14205/Desktop/npm/npm_test_data/v1"

	packageDir := filepath.Join(baseDir, packageName)
	if err := os.MkdirAll(packageDir, 0755); err != nil {
		log.Println("Failed to create package directory:", err)
		http.Error(w, "Failed to create directory", http.StatusInternalServerError)
		return
	}

	metaPath := filepath.Join(packageDir, "meta.json")
	metaFile, err := os.Create(metaPath)
	if err != nil {
		log.Println("Unable to create meta.json:", err)
		http.Error(w, "Unable to write metadata", http.StatusInternalServerError)
		return
	}

	encoder := json.NewEncoder(metaFile)
	encoder.SetIndent("", " ")
	if err := encoder.Encode(publishStruct); err != nil {
		metaFile.Close()
		http.Error(w, "Unable to encode metadata", http.StatusInternalServerError)
		return
	}
	metaFile.Close()
	log.Println("Metadata written to:", metaPath)

	for fileName, attachment := range publishStruct.Attachments {
		dataBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(attachment.Data))
		if err != nil {
			log.Println("Failed to decode base64 for", fileName, ":", err)
			continue
		}

		tarPath := filepath.Join(packageDir, fileName)
		file, err := os.Create(tarPath)
		if err != nil {
			log.Println("Unable to create file:", err)
			continue
		}

		if _, err := file.Write(dataBytes); err != nil {
			log.Println("Unable to write file:", err)
			file.Close()
			continue
		}

		if err := file.Close(); err != nil {
			log.Println("Error closing file:", err)
			continue
		}

		log.Printf("Saved attachment: %s (%d bytes)\n", tarPath, len(dataBytes))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"ok": true}`))
}
