package repositories

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"strings"
)

type ProductRepositoryInterface interface {
	FindAll() ([]models.ProductDomain, error)
	FindByID(id string) (*models.ProductDomain, error)
	Create(product models.ProductDomain) (*models.ProductDomain, error)
	Update(product models.ProductDomain) (*models.ProductDomain, error)
	Delete(id string) error
}

var PRODUCT_INDEX = "products"

type ProductRepository struct {
	client *elasticsearch.Client
	ctx    context.Context
}

func NewProductRepository(client *elasticsearch.Client, ctx context.Context) *ProductRepository {
	return &ProductRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *ProductRepository) FindAll() ([]models.ProductDomain, error) {
	res, err := r.client.Search(
		r.client.Search.WithContext(r.ctx),
		r.client.Search.WithIndex(PRODUCT_INDEX),
	)

	if err != nil {
		return nil, err
	}

	var products []models.ProductDomain
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		var product models.ProductDomain
		hitSource := hit.(map[string]interface{})["_source"]
		source, err := json.Marshal(hitSource)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(source, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *ProductRepository) FindByID(id string) (*models.ProductDomain, error) {
	res, err := r.client.Get(
		PRODUCT_INDEX,
		id,
		r.client.Get.WithContext(r.ctx),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, err
	}

	var product models.ProductDomain
	if err := json.NewDecoder(res.Body).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *ProductRepository) Create(product models.ProductDomain) (*models.ProductDomain, error) {
	id := uuid.New()
	product = models.NewProductDomain(id.String(), product.GetName(), product.GetPrice(), product.GetCategory())

	body, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Index(
		PRODUCT_INDEX,
		strings.NewReader(string(body)),
		r.client.Index.WithDocumentID(id.String()),
		r.client.Index.WithContext(r.ctx),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Update(product models.ProductDomain) (*models.ProductDomain, error) {
	product = models.NewProductDomain(product.GetID(), product.GetName(), product.GetPrice(), product.GetCategory())
	body, err := json.Marshal(product)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Update(
		PRODUCT_INDEX,
		product.GetID(),
		strings.NewReader(string(body)),
		r.client.Update.WithContext(r.ctx),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, err
	}

	return &product, nil
}

func (r *ProductRepository) Delete(id string) error {
	res, err := r.client.Delete(
		PRODUCT_INDEX,
		id,
		r.client.Delete.WithContext(r.ctx),
	)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return err
	}
	return nil
}
