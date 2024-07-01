package dtos

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type CategoryDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required,min=8,max=1000"`
}

func CategoryToDTO(category models.CategoryDomain) CategoryDTO {
	return CategoryDTO{
		ID:          category.GetID(),
		Name:        category.GetName(),
		Description: category.GetDescription(),
	}
}

func CategoryToDomain(category CategoryDTO) models.CategoryDomain {
	return models.NewCategoryDomain(category.ID, category.Name, category.Description)
}
