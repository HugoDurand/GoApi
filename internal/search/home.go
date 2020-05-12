package api

import(
	"net/http"
	"log"
	"io/ioutil"
)

func getBehanceContent()([]byte){
	//https://stats.nba.com/stats/boxscore
	res, err := http.Get("https://api.carbonintensity.org.uk/intensity")
	if err != nil {
		log.Fatal("Something went wrong getting Behance data")
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	print(contents)
        if err != nil {
            log.Fatal("Something went wrong getting Behance data")
        }
	return contents
}

type Response struct {
	Content []string
}

func ResponseHandler(w http.ResponseWriter, r *http.Request){
	//queryparams := r.URL.Query()["tags"]
	//responseContent := Response{queryparams}
	res:=getBehanceContent()
	//resp, error := json.Marshal(res)
	//if error != nil {
	//	http.Error(w, error.Error(), http.StatusInternalServerError)
	//	return
	//}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}