package main

import (
	"fmt"
	"gatewayDemo/combination"
	"gatewayDemo/gql"
	"gatewayDemo/guardian"
	"gatewayDemo/times"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	http.HandleFunc("/search/combination", combination.Search)
	http.HandleFunc("/search/times", times.Search)
	http.HandleFunc("/search/guardian", guardian.Search)
	http.HandleFunc("/search/graphql/guardian", gql.Search)

	fmt.Println("Now running on port 8090")
	http.ListenAndServe(":8090", nil)
}
