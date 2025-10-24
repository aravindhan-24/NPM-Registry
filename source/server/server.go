package server

import (
	"log"
	"net/http"
	"npm-registry/api"
	"npm-registry/utility"
	"strconv"

	"github.com/gorilla/mux"
)

var config *utility.Config

func init() {
	config = utility.ParseConfig()
	utility.CheckDir()
	db, err := utility.NewSQLiteDB(config.Data_dir.DB_File)
	if err != nil {
		log.Fatal("Unable to init sqllite db")
	}
	if !db.CreateTables() {
		log.Fatal("Unable to populate tables")
	}
	if !db.AddAdminCredentials() {
		log.Println("Unable to add admin credentials")
	}

}

func Serve() {
	r := mux.NewRouter()

	api.Register(r)

	log.Println("Npm registry started server at port :", config.Server.Port)
	if err := http.ListenAndServeTLS("localhost:"+strconv.Itoa(config.Server.Port),
		config.Server.Tls_config.Ca_cert,
		config.Server.Tls_config.Ca_key,
		r,
	); err != nil {
		log.Fatal("Unable to start server", err)
	}
}
