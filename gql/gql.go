package gql

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/graphql"
)

type Article struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	SectionID          string `json:"sectionId"`
	SectionName        string `json:"sectionName"`
	APIURL             string `json:"webPublicationDate"`
	WebURL             string `json:"webUrl"`
	WebTitle           string `json:"webTitle"`
	WebPublicationDate string `json:"apiUrl"`
	IsHosted           string `json:"isHosted"`
	PillarID           string `json:"pillarId"`
	PillarName         string `json:"pillarName"`
}

type Collection struct {
	Results []Article `json:"results"`
}

type ExtendedResponse struct {
	Response Collection `json:"response"`
}

func Search(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["subject"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	gKey := os.Getenv("GUARDIAN_KEY")
	response := makeGuardianCall(keys[0], gKey)

	var queryType = graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"ExtendedResponse": &graphql.Field{
					Type: ExtendedResponseType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return response, nil
					},
				},
			},
		},
	)

	var schema, _ = graphql.NewSchema(
		graphql.SchemaConfig{
			Query: queryType,
		},
	)
	result := executeQuery(req.URL.Query().Get("query"), schema)
	json.NewEncoder(w).Encode(result)

}

func makeGuardianCall(searchTerm string, key string) *ExtendedResponse {
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

func getResponse(body []byte) (*ExtendedResponse, error) {
	var s = new(ExtendedResponse)
	_ = json.Unmarshal(body, &s)
	return s, nil
}

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}
