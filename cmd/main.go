package cmd

import(
	"net/http"
	"api"
)

func main(){
	http.HandleFunc("/", api.homeHandler)
	http.ListenAndServe(":8080", nil)
}