package repositories

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type UserRepositoryInterface interface {
	FindAll() ([]models.UserDomain, error)
	FindByID(id string) (models.UserDomain, error)
	Create(user models.UserDomain) (models.UserDomain, error)
	Update(user models.UserDomain) (models.UserDomain, error)
	Delete(id int) error
}
