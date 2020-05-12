package handler

import(
	"net/http"
	"encoding/json"
	"../model"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	responseContent := model.Response{"Ma reponse"}
	resp, error := json.Marshal(responseContent)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}