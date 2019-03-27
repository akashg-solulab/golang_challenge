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

func FindProducts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var product Products
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	// if err := dao.Insert(movie); err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	fmt.Println("hello world",product)
	
	respondWithJson(w, http.StatusCreated, movie)
}

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
	fmt.Println("Dao",dao)
	dao.Connect()
}

// Define HTTP request routes
func main() {

	r := mux.NewRouter();
	r.HandleFunc("/products", FindProducts).Methods("POST")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
	
}
