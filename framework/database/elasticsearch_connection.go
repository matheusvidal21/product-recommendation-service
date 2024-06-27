package database

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	USER_INDEX          = "users"
	PRODUCT_INDEX       = "products"
	USER_ACTIVITY_INDEX = "user_activity"
	CATEGORY_INDEX      = "categories"
)

func NewElasticsearchConnection(address string) (*elasticsearch.Client, error) {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	res, err := client.Info()
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	indexing(client, USER_INDEX)
	indexing(client, PRODUCT_INDEX)
	indexing(client, USER_ACTIVITY_INDEX)
	indexing(client, CATEGORY_INDEX)

	logger.Info(fmt.Sprintf("Elasticsearch client: %s", res))
	return client, nil
}

func indexing(client *elasticsearch.Client, indexName string) {
	exists, err := indexExists(client, indexName)
	if err != nil {
		logger.Error("Error checking if index exists: %v", err)
		return
	}

	if !exists {
		err = createIndex(client, indexName)
		if err != nil {
			logger.Error("Error creating index: %v", err)
			return
		}
		logger.Info(fmt.Sprintf("Index %s created", indexName))
	} else {
		logger.Info(fmt.Sprintf("Index %s already exists", indexName))
	}
}

func indexExists(client *elasticsearch.Client, indexName string) (bool, error) {
	res, err := client.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}
	defer res.Body.Close()
	return res.StatusCode == 200, nil
}

func createIndex(client *elasticsearch.Client, indexName string) error {
	var mapping string
	switch indexName {
	case USER_INDEX:
		mapping = `{
				"mappings": {
					"properties": {
						"id": {"type": "keyword"},
						"name": {"type": "text"},
						"email": {"type": "keyword"},
						"password": {"type": "text"}
					}
				}
		}`
	case PRODUCT_INDEX:
		mapping = `{
				"mappings": {
					"properties": {
						"id": {"type": "keyword"},
						"name": {"type": "text"},
						"price": {"type": "double"},
						"category": {"type": "text"}
					}
				}
		}`
	case USER_ACTIVITY_INDEX:
		mapping = `{
				"mappings": {
					"properties": {
						"user_id": {"type": "keyword"},
						"product_id": {"type": "keyword"},
						"action": {"type": "text"}
					}
				}
		}`
	case CATEGORY_INDEX:
		mapping = `{
				"mappings": {
					"properties": {
						"id": {"type": "keyword"},	
						"name": {"type": "text"},	
						"description": {"type": "text"}
					}
				}
		}`
	}

	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(mapping),
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("Error creating index: %s", res.String())
	}
	return nil
}
