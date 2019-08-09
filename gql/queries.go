package gql

import "github.com/graphql-go/graphql"

var ArticleType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Article",
		Fields: graphql.Fields{
			"ID": &graphql.Field{
				Type: graphql.String,
			},
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"SectionID": &graphql.Field{
				Type: graphql.String,
			},
			"SectionName": &graphql.Field{
				Type: graphql.String,
			},
			"APIURL": &graphql.Field{
				Type: graphql.String,
			},
			"WebURL": &graphql.Field{
				Type: graphql.String,
			},
			"WebTitle": &graphql.Field{
				Type: graphql.String,
			},
			"WebPublicationDate": &graphql.Field{
				Type: graphql.String,
			},
			"IsHosted": &graphql.Field{
				Type: graphql.String,
			},
			"PillarID": &graphql.Field{
				Type: graphql.String,
			},
			"PillarName": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var CollectionType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Collection",
		Fields: graphql.Fields{
			"results": &graphql.Field{
				Type: graphql.NewList(ArticleType),
			},
		},
	},
)

var ExtendedResponseType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "ExtendedResponse",
		Fields: graphql.Fields{
			"response": &graphql.Field{
				Type: CollectionType,
			},
		},
	},
)
