package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/akashg-solulab/golang_challenge/config"
	. "github.com/akashg-solulab/golang_challenge/dao"
	. "github.com/akashg-solulab/golang_challenge/models"
)

var config = Config{}
var dao = ProductsDao{}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	// log.Fatal("hello")
	
	r := mux.NewRouter()
	
}
