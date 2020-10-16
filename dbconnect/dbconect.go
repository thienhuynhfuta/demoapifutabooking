package dbconnect

import (
	"bytes"
	"context"
	"encoding/json"
	"futaapibooking/common"
	"futaapibooking/model"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	es_Client    *elasticsearch.Client
	res_Client   *esapi.Response
	r            map[string]interface{}
	configclient model.ConfigElasticSearch
)

func InnitClientElasticsearch() {

	configclient = common.GetConfigElasticSearch()

	cfg := elasticsearch.Config{
		Addresses: []string{
			configclient.Host,
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	// 1. Get cluster info
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	// Check response status
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	// Deserialize the response into a map.
	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	es_Client = es
	res_Client = res
	// Print client and server version numbers.
	log.Printf("Client: %s && Server EasticSearch: %s", elasticsearch.Version, r["version"].(map[string]interface{})["number"])
}

func GetAreaSearch(query map[string]interface{}) ([]model.Area, error) {
	listArea := []model.Area{}
	var err error = nil
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
		return listArea, err
	}
	// Perform the search request.
	res_Client, err = es_Client.Search(
		es_Client.Search.WithContext(context.Background()),
		es_Client.Search.WithIndex(configclient.Index),
		es_Client.Search.WithBody(&buf),
		es_Client.Search.WithTrackTotalHits(true),
		es_Client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return listArea, err
	}
	defer res_Client.Body.Close()

	if res_Client.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res_Client.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res_Client.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
		return listArea, err
	}

	if err := json.NewDecoder(res_Client.Body).Decode(&r); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		result := hit.(map[string]interface{})["_source"]
		mydata, _ := json.Marshal(result)
		temp := model.Area{}
		json.Unmarshal([]byte(mydata), &temp)
		listArea = append(listArea, temp)
	}
	return listArea, err
}
