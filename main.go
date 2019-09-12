package main

import (
	"context"
	"log"
	"net/http"

	"github.com/ExPreman/go-elasticsearch/types"
	jsoniter "github.com/json-iterator/go"
	elastic "gopkg.in/olivere/elastic.v5"
)

var esClient *elastic.Client

func init() {
	httpClient := &http.Client{Transport: &types.DefaultHeaderTransport{}}

	var err error
	esClient, err = elastic.NewClient(
		elastic.SetURL("http://127.0.0.1:9200"),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
		elastic.SetDecoder(&types.Decoder{}),
		elastic.SetHttpClient(httpClient),
	)
	if err != nil {
		log.Fatal(err)
	}
}

type Product struct {
	CategoryRecommendation struct {
		ID   int64  `json:"id"`
		Name string `json:"name"`
	} `json:"category_recommendation"`

	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Score int64  `json:"score"`
}

func main() {
	// GET BY ID
	resGet, err := esClient.Get().Index("autocomplete_v1").Id("1").Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	if resGet.Source == nil {
		log.Fatal("empty source")
	}

	var data Product
	err = jsoniter.ConfigFastest.Unmarshal(*resGet.Source, &data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("GET Data Result: ", data)

	// SEARCH
	query := elastic.NewTermQuery("name", "samsung")
	resSearch, err := esClient.Search().
		Index("autocomplete_v1").
		Type("_doc").
		Query(query).
		Sort("score", false).
		From(0).Size(10).
		Pretty(true).
		ReadSource().
		Do(context.TODO())
	if err != nil {
		log.Fatal("error query: ", err)
	}

	var dataSearch []Product
	if resSearch != nil {
		if resSearch.Hits == nil || resSearch.Hits.Hits == nil || len(resSearch.Hits.Hits) == 0 {
			log.Fatal("data not found")
		}
		for _, hit := range resSearch.Hits.Hits {
			if hit.Source == nil {
				continue
			}

			var tmp Product
			err = jsoniter.ConfigFastest.Unmarshal(*hit.Source, &tmp)
			if err != nil {
				continue
			}
			dataSearch = append(dataSearch, tmp)
		}
	}
	log.Println("GET Data Search: ", dataSearch)
}
