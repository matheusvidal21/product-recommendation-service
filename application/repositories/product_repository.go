package repositories

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type ProductRepositoryInterface interface {
	FindAll() ([]models.ProductDomain, error)
	FindByID(id string) (models.ProductDomain, error)
	Create(product models.ProductDomain) (models.ProductDomain, error)
	Update(product models.ProductDomain) (models.ProductDomain, error)
	Delete(id int) error
}
