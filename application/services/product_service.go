package services

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
)

type ProductServiceInterface interface {
	GetAllProducts() ([]dtos.ProductDTO, error)
	GetProductByID(id string) (*models.ProductDomain, error)
	CreateProduct(name string, price float64, categoryId string) (*models.ProductDomain, error)
	UpdateProduct(id, name string, price float64, categoryId string) (*models.ProductDomain, error)
	DeleteProduct(id string) error
}

type ProductService struct {
	categoryService CategoryServiceInterface
	repo            repositories.ProductRepositoryInterface
	ctx             context.Context
}

func NewProductService(repo repositories.ProductRepositoryInterface, categoryService CategoryServiceInterface, ctx context.Context) ProductServiceInterface {
	return &ProductService{
		repo:            repo,
		categoryService: categoryService,
		ctx:             ctx,
	}
}

func (s *ProductService) GetAllProducts() ([]dtos.ProductDTO, error) {
	logger.Info("Getting all products services")
	prods := []dtos.ProductDTO{}
	products, err := s.repo.FindAll()
	if err != nil {
		return prods, err
	}

	for _, p := range products {
		prods = append(prods, dtos.ProductToDTO(p))
	}

	return prods, nil
}

func (s *ProductService) GetProductByID(id string) (*models.ProductDomain, error) {
	logger.Info("Getting product by id services")
	return s.repo.FindByID(id)
}

func (s *ProductService) CreateProduct(name string, price float64, categoryId string) (*models.ProductDomain, error) {
	logger.Info("Creating product services")

	cat, err := s.categoryService.GetCategoryByID(categoryId)
	if err != nil || cat == nil {
		logger.Error("Error getting category", err)
		return nil, err
	}

	id := uuid.New().String()
	product := models.NewProductDomain(id, name, price, models.NewCategoryDomain(categoryId, "", ""))
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(id, name string, price float64, categoryId string) (*models.ProductDomain, error) {
	logger.Info("Updating product services")

	_, err := s.categoryService.GetCategoryByID(categoryId)
	if err != nil {
		logger.Error("Error getting category", err)
		return nil, err
	}

	product := models.NewProductDomain(id, name, price, models.NewCategoryDomain(categoryId, "", ""))
	return s.repo.Update(id, product)
}

func (s *ProductService) DeleteProduct(id string) error {
	logger.Info("Deleting product services")
	return s.repo.Delete(id)
}
