package guardian

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
	WebURL             string
	WebTitle           string
	WebPublicationDate string
}

type Collection struct {
	Results []Article `json:"results"`
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

	gKey := os.Getenv("GUARDIAN_KEY")
	response := makeGuardianCall(keys[0], gKey)

	b, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(b)

}

func InternalSearch(searchTerm string) *Response {
	gKey := os.Getenv("GUARDIAN_KEY")
	return makeGuardianCall(searchTerm, gKey)
}

func makeGuardianCall(searchTerm string, key string) *Response {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	gURL := os.Getenv("GUARDIAN_URL")
	fullURL := fmt.Sprintf(gURL, searchTerm, key)

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
