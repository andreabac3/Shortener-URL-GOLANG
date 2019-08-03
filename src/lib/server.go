package lib

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var router = mux.NewRouter()
var Client, collection = W_GetCollection()

func Start_Server() int8 {

	router.HandleFunc("/insert/{url}", InsertURL_Server)
	router.HandleFunc("/get/{url}", GetURL_Server)
	router.HandleFunc("/get_no_redirect/{url}", GetURL_No_Redirect_Server)

	http.Handle("/", router)
	port := os.Getenv("SERVER_PORT")
	port = ":" + port

	if http.ListenAndServe(port, router) != nil {
		return -1
	}
	return 0
}
func InsertURL_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result := InsertURL(collection, vars["url"])
	fmt.Fprintf(w, "Hash: %v\n", result)
}

func GetURL_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result := QueryURL(collection, vars["url"])
	result = "http://" + result
	http.Redirect(w, r, result, http.StatusSeeOther)
	fmt.Println(result)
}

func GetURL_No_Redirect_Server(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["url"])
	result := QueryURL(collection, vars["url"])
	result = "http://" + result
	fmt.Fprintf(w, "Original: %v\n", result)
	fmt.Println(result)
}

