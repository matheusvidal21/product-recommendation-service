package repositories

import (
	"context"
	"database/sql"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/framework/database"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
)

type CategoryRepositoryInterface interface {
	FindAll() ([]models.CategoryDomain, error)
	FindByID(id string) (*models.CategoryDomain, error)
	Create(category models.CategoryDomain) (*models.CategoryDomain, error)
	Update(id string, category models.CategoryDomain) (*models.CategoryDomain, error)
	Delete(id string) error
}

type CategoryRepository struct {
	queries *database.Queries
	ctx     context.Context
}

func NewCategoryRepository(db *sql.DB, ctx context.Context) CategoryRepositoryInterface {
	return &CategoryRepository{
		queries: database.New(db),
		ctx:     ctx,
	}
}

func (r *CategoryRepository) FindAll() ([]models.CategoryDomain, error) {
	logger.Info("Fetching all categories")
	categories, err := r.queries.GetAllCategories(r.ctx)
	if err != nil {
		logger.Error("Error fetching categories", err)
		return nil, err
	}

	var categoryDomains []models.CategoryDomain
	for _, category := range categories {
		categoryDomain := models.NewCategoryDomain(category.ID, category.Name, category.Description.String)
		categoryDomains = append(categoryDomains, categoryDomain)
	}

	logger.Info("Categories fetched")
	return categoryDomains, nil
}

func (r *CategoryRepository) FindByID(id string) (*models.CategoryDomain, error) {
	logger.Info("Fetching category by ID")
	category, err := r.queries.GetCategoryByID(r.ctx, id)
	if err != nil {
		logger.Error("Error fetching category", err)
		return nil, err
	}

	categoryDomain := models.NewCategoryDomain(category.ID, category.Name, category.Description.String)
	return &categoryDomain, nil
}

func (r *CategoryRepository) Create(category models.CategoryDomain) (*models.CategoryDomain, error) {
	logger.Info("Creating category")
	newCategory, err := r.queries.CreateCategory(r.ctx, database.CreateCategoryParams{
		ID:          category.GetID(),
		Name:        category.GetName(),
		Description: sql.NullString{String: category.GetDescription(), Valid: true},
	})

	if err != nil {
		logger.Error("Error creating category", err)
		return nil, err
	}

	logger.Info("Category created")
	category = models.NewCategoryDomain(newCategory.ID, newCategory.Name, newCategory.Description.String)
	return &category, nil
}

func (r *CategoryRepository) Update(id string, category models.CategoryDomain) (*models.CategoryDomain, error) {
	logger.Info("Updating category")
	category, err := r.queries.UpdateCategory(r.ctx, database.UpdateCategoryParams{
		ID:          id,
		Name:        category.GetName(),
		Description: sql.NullString{String: category.GetDescription(), Valid: true},
	})

	if err != nil {
		logger.Error("Error updating category", err)
		return nil, err
	}

	category = models.NewCategoryDomain(id, category.GetName(), category.GetDescription())
	logger.Info("Category updated")
	return &category, nil
}

func (r *CategoryRepository) Delete(id string) error {
	logger.Info("Deleting category")
	err := r.queries.DeleteCategory(r.ctx, id)
	if err != nil {
		logger.Error("Error deleting category", err)
		return err
	}

	logger.Info("Category deleted")
	return nil
}
