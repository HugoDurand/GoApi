package api

import(
	"net/http"
	"log"
	"io/ioutil"
	"time"
)

func getNotWorkingFirstContent(c chan []byte){
	client := http.Client{
		Timeout: 3 * time.Second,
	}
	res, err := client.Get("https://stats.nba.com/stats/boxscore")
	if err != nil {
		close(c)
		log.Print("Something went wrong getting other data")
		return
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	print("First data")
        if err != nil {
            close(c)
			log.Print("Something went wrong reading other data")
			return
        }
	c <- contents
}

func getSecondContent(c chan []byte){
	res, err := http.Get("https://restcountries.eu/rest/v2/all")
	if err != nil {
		close(c)
		log.Print("Something went wrong getting other data")
		return
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	print("Second data")
        if err != nil {
            close(c)
			log.Print("Something went wrong reading other data")
			return
        }
	c <- contents
}

func getThirdContent(c chan []byte){
	res, err := http.Get("https://api.carbonintensity.org.uk/intensity")
	if err != nil {
		close(c)
		log.Print("Something went wrong getting Behance data")
		return
	}
	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	print("Third data")
        if err != nil {
			close(c)
			log.Print("Something went wrong reading Behance data")
			return
        }
	c <- contents
}

func ResponseHandler(w http.ResponseWriter, r *http.Request){

	c:= make(chan []byte)
	go getNotWorkingFirstContent(c)
	go getSecondContent(c)
	go getThirdContent(c)
	res, res2, res3 := <-c, <-c, <-c

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write(res2)
	w.Write(res3)
}