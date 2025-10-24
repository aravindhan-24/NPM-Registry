package api

import (
	"npm-registry/handler"

	"github.com/gorilla/mux"
)

func Register(r *mux.Router) *mux.Router {
	r.HandleFunc("/v1", handler.PingHandler).Methods("GET")

	r.HandleFunc("/-/user/org.couchdb.user:{name:.*}", handler.AuthorizeUser).Methods("PUT")
	r.HandleFunc("/{packagename:.+}", handler.HandlePublish).Methods("PUT")

	r.HandleFunc("/{packagename:.+}/-/{packagetar:[a-zA-Z0-9_.-]+\\.tgz}", handler.ServeTarBall).Methods("GET")
	r.HandleFunc("/{packagename:.+}", handler.ServeMetaData).Methods("GET")

	return r

}
