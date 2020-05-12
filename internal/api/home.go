package api

import(
	"net/http"
	"encoding/json"
)

type Response struct {
	Content string
}

func responseHandler(w http.ResponseWriter, r *http.Request){
	responseContent := Response{"Ma reponse"}
	resp, error := json.Marshal(responseContent)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}