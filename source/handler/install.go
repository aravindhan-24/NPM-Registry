package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ServeTarBall(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authorization header", r.Header.Get("Authorization"))
}

func ServeMetaData(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Authorization header", r.Header.Get("Authorization"))
	metadata := map[string]interface{}{
		"name": "mypackage",
		"dist-tags": map[string]string{
			"latest": "1.0.0",
		},
		"versions": map[string]interface{}{
			"1.0.0": map[string]interface{}{
				"name":    "mypackage",
				"version": "1.0.0",
				"dist": map[string]string{
					"shasum":    "d5b91466b0b6752f224221361af54c4cece28883198a80e2d1c03ec456d6800d506d529a0a78bfe46bfd28b126bda1ed9204ab0ac4153a084b2795d8b997b687",
					"tarball":   "http://localhost:2424/npm-test-temp-1.0.0/-/npm-test-temp-1.0.0.tgz",
					"integrity": "sha512-BASE64_SHA512_HASH",
				},
				"dependencies": map[string]string{},
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metadata)
}
