package repositories

import (
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/olivere/elastic/v7"
)

type UserRepositoryInterface interface {
	FindAll() ([]models.UserDomain, error)
	FindByID(id string) (models.UserDomain, error)
	Create(user models.UserDomain) (models.UserDomain, error)
	Update(user models.UserDomain) (models.UserDomain, error)
	Delete(id int) error
}

type UserRepository struct {
}

func NewUserRepository(client *elastic.Client, index string) *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindAll() ([]models.UserDomain, error) {
	return nil, nil
}

func (r *UserRepository) FindByID(id string) (models.UserDomain, error) {
	return nil, nil

}

func (r *UserRepository) Create(user models.UserDomain) (models.UserDomain, error) {
	return nil, nil

}

func (r *UserRepository) Update(user models.UserDomain) (models.UserDomain, error) {
	return nil, nil
}

func (r *UserRepository) Delete(id int) error {
	return nil
}
