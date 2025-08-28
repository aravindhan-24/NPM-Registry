package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"npm-registry/constants"
	"os"
	"strings"
)

func HandlePublish(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authorization header", r.Header.Get("Authorization"))
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Unable to read body ", err)
	}
	var publishStruct constants.Package
	if err := json.Unmarshal(body, &publishStruct); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}
	for fileName, attachment := range publishStruct.Attachments {
		dataBytes, err := base64.StdEncoding.DecodeString(strings.TrimSpace(attachment.Data))
		if err != nil {
			fmt.Println("Failed to decode base64 for", fileName, ":", err)
			continue
		}

		fmt.Println("fileName ", fileName)
		// read root path from config and write in that path
		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Unable to create file")
			continue
		}
		count, err := file.Write(dataBytes)
		if count != attachment.Length {
			http.Error(w, fmt.Sprintf("Size mismatch for %s: expected %d, got %d", fileName, attachment.Length, len(dataBytes)), http.StatusBadRequest)
			return
		}
		if err != nil {
			fmt.Println("unable to write file:", err)
			file.Close()
			continue
		}
		if err := file.Close(); err != nil {
			fmt.Println("error closing file:", err)
			continue
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"ok": true}`))
}
