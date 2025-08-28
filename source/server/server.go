package server

import (
	"fmt"
	"net/http"
	"npm-registry/api"

	"github.com/gorilla/mux"
)

func Serve() {
	r := mux.NewRouter()

	api.Register(r)

	fmt.Println("Npm registry started server at port :2424")
	if err := http.ListenAndServe("localhost:2424", r); err != nil {
		fmt.Println("Unable to start server")
	}
}
