package database

import (
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
	"net"
	"net/http"
	"time"
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
	logger.Info(fmt.Sprintf("Elasticsearch client: %s", res))
	return client, nil
}
