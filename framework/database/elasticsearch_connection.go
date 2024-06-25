package database

import (
	"context"
	"fmt"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
	"github.com/olivere/elastic/v7"
)

func NewElasticsearchConnection(address string) (*elastic.Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(address), elastic.SetSniff(false))

	if err != nil {
		return nil, err
	}

	info, code, err := client.Ping(address).Do(context.Background())
	if err != nil {
		return nil, err
	}
	logger.Info(fmt.Sprintf("Elasticsearch returned with code %d and version %s", code, info.Version.Number))

	return client, nil
}
