package api

import(
	"net/http"
	"./handler"
)

func main(){
	http.HandleFunc("/", handler.homeHandler)
	http.ListenAndServe(":8080", nil)
}