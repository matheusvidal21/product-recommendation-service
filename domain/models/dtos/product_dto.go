package dtos

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type ProductDTO struct {
	ID       string      `json:"id"`
	Name     string      `json:"name" validate:"required,min=3,max=100"`
	Price    float64     `json:"price" validate:"required,gt=0"`
	Category CategoryDTO `json:"category" validate:"required"`
}

func ProductToDTO(product models.ProductDomain) ProductDTO {
	return ProductDTO{
		ID:    product.GetID(),
		Name:  product.GetName(),
		Price: product.GetPrice(),
	}
}

func ProductToDomain(product ProductDTO) models.ProductDomain {
	return models.NewProductWithId(product.ID, product.Name, product.Price, CategoryToDomain(product.Category))
}
