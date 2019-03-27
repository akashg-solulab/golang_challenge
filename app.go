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
var dao = ProductsDAO{}


func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
func FindSuggestionEndpoint(w http.ResponseWriter, r *http.Request) {

	var cm Coffee_Machine
	fmt.Printf("paramsparams: %s",r.Body)


	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&cm); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	cm.ID = bson.NewObjectId()

		fmt.Printf("paramsparams: %s",cm.Product)

	

	// params := mux.Vars(r)
	// body := mux.Vars(r)

	// fmt.Printf("paramsparams%s%s",params["id"],"abcd")

	// key:= params["id"]

	// fmt.Printf("paramsparams: %s/n",key)

	

	// myArr := strings.Split(key, " ")
	// // for index := 0; index < len(myArr); index++ {
	// 	fmt.Println("muArr ",myArr[0])	
	// }

	// product
	// product_type
	// coffe_flavor
	// pack_size

	// if (len(myvar) == 1){

	// 	machine - query all with machine keyword
	// 	// flavor - query all with flavor
	// 	// machine size - query all with machine size that is given
	// 	// packsize - query all with packsize
	// 	// single_query(type,value)	

	// }
	// else if(len(myvar) == 2){

	// 	product & product type- machine, small : query all machines with both the arguments	
	// 											query again with product as pod and product_type as small
	// 	product and flavor - machine, flavor	all machines and all flavors ?
	// 						 prod, flavor		all prods and flavor 
	// 	product and packsize - machine, packsize	all machines
	// 							prod, packsize	all prods and pack_size\
		

		
	// }
	// else if(len(myvar) == 3){

	// 	product ,product_type and flavor - machine , small, vanilla
	// 										prod, small, vanilla
		
	// 	product ,product_type, pack size - machine, small, 3 - query with machines and 
	// 										prod, small, 3  
					
									
	// }
	// else if(len(myvar) == 1){

		
	// }
	// else {
		
	// 	fmt.Println("Invalid no of arguments")
	// }


	// movie, err := dao.FindById(params["id"])
	// if err != nil {

	// 	respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
	// 	return
	// }
	// respondWithJson(w, http.StatusOK, movie)

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
	r.HandleFunc("/task", FindSuggestionEndpoint).Methods("GET")
	
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
