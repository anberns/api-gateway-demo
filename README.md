# Demo API Gateway
This project is an attempt to make a simple API Gateway as proposed in "Microservices Patterns" by Chris Richardson. It functions as an intermediary between front-end applications, external consumers, and internal services, allowing for both the selection from, and the combining of, microservice data without changing what the microservice chooses to make available through its own API. 

This sample Gateway draws data from two source APIs, the NY Times and the Guardian, which in this example represent two separate microservices. It offers standard `GET` routes for searching for articles from either news organization separately, as well as a route that combines or 'stitches' and sorts the data from both sources for a given search term. A simple GraphQL endpoint is also included that exposes the full data of the Guardian API and allows for custom attribute selection.

## Dev Setup

To run locally, run `go run main.go` from the project's root directory. The local api server will be available at `localhost:8090`.

### Semi-RESTful Endpoints
`/search/times/?subject={subject}` - Exposes selected data drawn from the NY Times article search api.

`GET` - Receives a search subject as a URL parameter. Returns a JSON object containing an array of articles with fields `Web_URL`, `Headline`, and `Pub_Date`.


`/search/guardian/?subject={subject}` - Exposes selected data drawn from the Guardian article search api.

`GET` - Receives a search subject as a URL parameter. Returns a JSON object containing an array of articles with fields `WebURL`, `WebTitle`, and `WebPublicationDate`.


`/search/combination/?subject={subject}` - Exposes selected data drawn from both the NY Times and the Guardian article search apis. Ten articles are taken from each source, combined, and sorted by descending publication date.

`GET` - Receives a search subject as a URL parameter. Returns a JSON object containing an array of articles with fields `URL`, `Title`, and `Date`.

### GraphQL Endpoint

`/search/graphql/guardian/?search={subject}&query={query}` - Exposes the full data available from the Guardian api while offering the ability to choose any combination of fields through a GraphQL `Query`.

`GET` - Receives a search subject as a URL parameter as well a GraphQL query in the form. Returns a JSON object containing an array of articles with the chosen fields.

Available Fields: `ID`, `Type`, `SectionID`, `SectionName`, `APIURL`, `WebURL`, `WebTitle`, `WebPublicationDate`, `IsHosted`, `PillarID`, and `PillarName`.


##### Example Call
Returns a set of articles with the subject of China and the fields `WebTitle` and `WebURL`.
```
localhost:8090/search/graphql/guardian?subject=china&query={ExtendedResponse{response{results{WebTitle, WebURL}}}}
```

