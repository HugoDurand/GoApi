package main

import(
	"net/http"
	"github.com/HugoDurand/MyDesignBoard/internal/api"
)

func main(){
	http.HandleFunc("/", api.ResponseHandler)
	http.ListenAndServe(":8080", nil)
}