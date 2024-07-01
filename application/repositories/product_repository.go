package repositories

import (
	"context"
	"database/sql"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
	"github.com/matheusvidal21/product-recommendation-service/framework/database"
)

type ProductRepositoryInterface interface {
	FindAll() ([]models.ProductDomain, error)
	FindByID(id string) (*models.ProductDomain, error)
	Create(product models.ProductDomain) (*models.ProductDomain, error)
	Update(id string, product models.ProductDomain) (*models.ProductDomain, error)
	Delete(id string) error
}

type ProductRepository struct {
	queries *database.Queries
	ctx     context.Context
}

func NewProductRepository(db *sql.DB, ctx context.Context) ProductRepositoryInterface {
	return &ProductRepository{
		queries: database.New(db),
		ctx:     ctx,
	}
}

func (r *ProductRepository) FindAll() ([]models.ProductDomain, error) {
	logger.Info("Fetching all products")
	products, err := r.queries.GetAllProducts(r.ctx)
	if err != nil {
		logger.Error("Error fetching products", err)
		return nil, err
	}

	var productDomains []models.ProductDomain
	for _, product := range products {
		category := models.NewCategoryDomain(product.CategoryID.String, "", "")
		productDomain := models.NewProductDomain(product.ID, product.Name, product.Price, category)
		productDomains = append(productDomains, productDomain)
	}

	logger.Info("Products fetched")
	return productDomains, nil
}

func (r *ProductRepository) FindByID(id string) (*models.ProductDomain, error) {
	logger.Info("Fetching product by ID")
	product, err := r.queries.GetProductByID(r.ctx, id)
	if err != nil {
		logger.Error("Error fetching product", err)
		return nil, err
	}
	category := models.NewCategoryDomain(product.CategoryID.String, "", "")
	productDomain := models.NewProductDomain(product.ID, product.Name, product.Price, category)
	return &productDomain, nil
}

func (r *ProductRepository) Create(product models.ProductDomain) (*models.ProductDomain, error) {
	logger.Info("Creating product")
	var newProduct database.Product
	var err error

	if product.GetCategory().GetID() != "" {
		newProduct, err = r.queries.CreateProduct(r.ctx, database.CreateProductParams{
			ID:         product.GetID(),
			Name:       product.GetName(),
			Price:      product.GetPrice(),
			CategoryID: sql.NullString{String: product.GetCategory().GetID(), Valid: true},
		})
	} else {
		newProduct, err = r.queries.CreateProduct(r.ctx, database.CreateProductParams{
			ID:    product.GetID(),
			Name:  product.GetName(),
			Price: product.GetPrice(),
		})
	}

	if err != nil {
		logger.Error("Error creating product", err)
		return nil, err
	}

	logger.Info("Product created")
	product = models.NewProductDomain(newProduct.ID, newProduct.Name, newProduct.Price, product.GetCategory())
	return &product, nil
}

func (r *ProductRepository) Update(id string, product models.ProductDomain) (*models.ProductDomain, error) {
	logger.Info("Updating product")
	var updatedProduct database.Product
	var err error

	if product.GetCategory().GetID() != "" {
		updatedProduct, err = r.queries.UpdateProduct(r.ctx, database.UpdateProductParams{
			ID:         id,
			Name:       product.GetName(),
			Price:      product.GetPrice(),
			CategoryID: sql.NullString{String: product.GetCategory().GetID(), Valid: true},
		})
	} else {
		updatedProduct, err = r.queries.UpdateProduct(r.ctx, database.UpdateProductParams{
			ID:    id,
			Name:  product.GetName(),
			Price: product.GetPrice(),
		})
	}

	if err != nil {
		logger.Error("Error updating product", err)
		return nil, err
	}

	product = models.NewProductDomain(updatedProduct.ID, updatedProduct.Name, updatedProduct.Price, product.GetCategory())
	logger.Info("Product updated")
	return &product, nil
}

func (r *ProductRepository) Delete(id string) error {
	logger.Info("Deleting product")
	err := r.queries.DeleteProduct(r.ctx, id)
	if err != nil {
		logger.Error("Error deleting product", err)
		return err
	}

	logger.Info("Product deleted")
	return nil
}
