package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/services"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	"net/http"
)

type ProductControllerInterface interface {
	Save(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type ProductController struct {
	productService services.ProductServiceInterface
}

func NewProductController(productService services.ProductServiceInterface) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) Save(c *fiber.Ctx) error {
	var productDTO dtos.ProductDTO
	if err := c.BodyParser(&productDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(productDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := pc.productService.CreateProduct(productDTO.Name, productDTO.Price, dtos.CategoryToDomain(productDTO.Category))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

func (pc *ProductController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	product, err := pc.productService.GetProductByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

func (pc *ProductController) GetAll(c *fiber.Ctx) error {
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(products)
}

func (pc *ProductController) Update(c *fiber.Ctx) error {
	var productDTO dtos.ProductDTO
	if err := c.BodyParser(&productDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(productDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	product, err := pc.productService.UpdateProduct(productDTO.ID, productDTO.Name, productDTO.Price, dtos.CategoryToDomain(productDTO.Category))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

func (pc *ProductController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	err = pc.productService.DeleteProduct(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
