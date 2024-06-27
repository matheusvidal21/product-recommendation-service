package dtos

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type ProductDTO struct {
	ID       string      `json:"id"`
	Name     string      `json:"name"`
	Price    float64     `json:"price"`
	Category CategoryDTO `json:"category"`
}

func ProductToDTO(product models.ProductDomain) ProductDTO {
	return ProductDTO{
		ID:    product.GetID(),
		Name:  product.GetName(),
		Price: product.GetPrice(),
	}
}

func ProductToDomain(product ProductDTO) models.ProductDomain {
	return models.NewProductDomain(product.ID, product.Name, product.Price, CategoryToDomain(product.Category))
}
