package times

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Article struct {
	Web_URL  string
	Headline Headline
	Pub_Date string
}

type Headline struct {
	Main string
}

type Collection struct {
	Docs []Article `json:"docs"`
}

type Response struct {
	Response Collection `json:"response"`
}

func Search(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["subject"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	tKey := os.Getenv("TIMES_KEY")
	response := makeTimesCall(keys[0], tKey)

	b, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func InternalSearch(searchTerm string) *Response {
	tKey := os.Getenv("TIMES_KEY")
	return makeTimesCall(searchTerm, tKey)
}

func makeTimesCall(searchTerm string, key string) *Response {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	tURL := os.Getenv("TIMES_URL")
	fullURL := fmt.Sprintf(tURL, searchTerm, key)

	request, _ := http.NewRequest("GET", fullURL, nil)
	res, _ := client.Do(request)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	response, _ := getResponse([]byte(body))
	return response
}

func getResponse(body []byte) (*Response, error) {
	var s = new(Response)
	_ = json.Unmarshal(body, &s)
	return s, nil
}
