package repositories

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"strings"
)

var USER_INDEX = "users"

type UserRepositoryInterface interface {
	FindAll() ([]models.UserDomain, error)
	FindByID(id string) (*models.UserDomain, error)
	Create(user models.UserDomain) (*models.UserDomain, error)
	Update(user models.UserDomain) (*models.UserDomain, error)
	Delete(id string) error
}

type UserRepository struct {
	client *elasticsearch.Client
	ctx    context.Context
}

func NewUserRepository(client *elasticsearch.Client, ctx context.Context) *UserRepository {
	return &UserRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *UserRepository) FindAll() ([]models.UserDomain, error) {
	res, err := r.client.Search(
		r.client.Search.WithContext(r.ctx),
		r.client.Search.WithIndex(USER_INDEX),
	)
	if err != nil {
		return nil, err
	}

	var users []models.UserDomain
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		var user models.UserDomain
		hitSource := hit.(map[string]interface{})["_source"]
		source, err := json.Marshal(hitSource)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(source, &user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) FindByID(id string) (*models.UserDomain, error) {
	res, err := r.client.Get(
		USER_INDEX,
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

	var user models.UserDomain
	if err := json.NewDecoder(res.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(user models.UserDomain) (*models.UserDomain, error) {
	id := uuid.New()
	user = models.NewUser(id.String(), user.GetName(), user.GetEmail(), user.GetPassword())

	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Index(
		USER_INDEX,
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

	return &user, nil
}

func (r *UserRepository) Update(user models.UserDomain) (*models.UserDomain, error) {
	user = models.NewUser(user.GetID(), user.GetName(), user.GetEmail(), user.GetPassword())
	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Update(
		USER_INDEX,
		user.GetID(),
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

	return &user, nil
}

func (r *UserRepository) Delete(id string) error {
	res, err := r.client.Delete(
		USER_INDEX,
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
