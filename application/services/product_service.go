package services

import (
	"context"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
)

type ProductServiceInterface interface {
	GetAllProducts() ([]models.ProductDomain, error)
	GetProductByID(id string) (*models.ProductDomain, error)
	CreateProduct(name string, price float64, category models.CategoryDomain) (*models.ProductDomain, error)
	UpdateProduct(id, name string, price float64, category models.CategoryDomain) (*models.ProductDomain, error)
	DeleteProduct(id string) error
}

type ProductService struct {
	repo repositories.ProductRepositoryInterface
	ctx  context.Context
}

func NewProductService(repo repositories.ProductRepositoryInterface, ctx context.Context) ProductServiceInterface {
	return &ProductService{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *ProductService) GetAllProducts() ([]models.ProductDomain, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductByID(id string) (*models.ProductDomain, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) CreateProduct(name string, price float64, category models.CategoryDomain) (*models.ProductDomain, error) {
	product := models.NewProductDomain("", name, price, category)
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(id, name string, price float64, category models.CategoryDomain) (*models.ProductDomain, error) {
	product := models.NewProductDomain(id, name, price, category)
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.repo.Delete(id)
}
