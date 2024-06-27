package repositories

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"strings"
)

var CATEGORY_INDEX = "categories"

type CategoryRepositoryInterface interface {
	FindAll() ([]models.CategoryDomain, error)
	FindByID(id string) (*models.CategoryDomain, error)
	Create(category models.CategoryDomain) (*models.CategoryDomain, error)
	Update(category models.CategoryDomain) (*models.CategoryDomain, error)
	Delete(id string) error
}

type CategoryRepository struct {
	client *elasticsearch.Client
	ctx    context.Context
}

func NewCategoryRepository(client *elasticsearch.Client, ctx context.Context) *CategoryRepository {
	return &CategoryRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *CategoryRepository) FindAll() ([]models.CategoryDomain, error) {
	res, err := r.client.Search(
		r.client.Search.WithContext(r.ctx),
		r.client.Search.WithIndex(CATEGORY_INDEX),
	)
	if err != nil {
		return nil, err
	}

	var categories []models.CategoryDomain
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		var category models.CategoryDomain
		hitSource := hit.(map[string]interface{})["_source"]
		source, err := json.Marshal(hitSource)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(source, &category)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (r *CategoryRepository) FindByID(id string) (*models.CategoryDomain, error) {
	res, err := r.client.Get(
		CATEGORY_INDEX,
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

	var category models.CategoryDomain
	if err := json.NewDecoder(res.Body).Decode(&category); err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *CategoryRepository) Create(category models.CategoryDomain) (*models.CategoryDomain, error) {
	id := uuid.New()
	category = models.NewCategoryDomain(id.String(), category.GetName(), category.GetDescription())

	body, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Index(
		CATEGORY_INDEX,
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

	return &category, nil
}

func (r *CategoryRepository) Update(category models.CategoryDomain) (*models.CategoryDomain, error) {
	category = models.NewCategoryDomain(category.GetID(), category.GetName(), category.GetDescription())
	body, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Update(
		CATEGORY_INDEX,
		category.GetID(),
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

	return &category, nil
}

func (r *CategoryRepository) Delete(id string) error {
	res, err := r.client.Delete(
		CATEGORY_INDEX,
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
