package combination

import (
	"encoding/json"
	"gatewayDemo/guardian"
	"gatewayDemo/times"
	"log"
	"net/http"
	"sort"
	"sync"
)

type Collection struct {
	Articles []Article
}

type Article struct {
	URL   string
	Title string
	Date  string
}

func Search(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["subject"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	searchTerm := keys[0]

	var timesData = new(times.Response)
	var guardianData = new(guardian.Response)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		timesData = times.InternalSearch(searchTerm)
	}()
	go func() {
		defer wg.Done()
		guardianData = guardian.InternalSearch(searchTerm)
	}()

	wg.Wait()

	collection := stitch(timesData, guardianData)
	c, _ := json.Marshal(collection)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(c)
}

func stitch(tData *times.Response, gData *guardian.Response) *Collection {

	var c = new(Collection)
	var a = new(Article)

	tArticles := tData.Response.Docs
	for i, _ := range tArticles {
		a.URL = tArticles[i].Web_URL
		a.Title = tArticles[i].Headline.Main
		a.Date = tArticles[i].Pub_Date
		c.Articles = append(c.Articles, *a)
	}

	gArticles := gData.Response.Results
	for i, _ := range gArticles {
		a.URL = gArticles[i].WebURL
		a.Title = gArticles[i].WebTitle
		a.Date = gArticles[i].WebPublicationDate
		c.Articles = append(c.Articles, *a)
	}

	cSortDesc(c)
	return c
}

func cSortDesc(c *Collection) {
	sort.Slice(c.Articles, func(i, j int) bool {
		return c.Articles[i].Date > c.Articles[j].Date
	})
}
