package handlers

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/services"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
	"github.com/matheusvidal21/product-recommendation-service/framework/config/rest_err"
	"github.com/matheusvidal21/product-recommendation-service/framework/config/validation"
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
	logger.Info("Saving product")
	var productDTO dtos.ProductDTO
	if err := c.BodyParser(&productDTO); err != nil {
		logger.Error("Error parsing body", err)
		restErr := rest_err.NewBadRequestError("invalid json body")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(productDTO); err != nil {
		logger.Error("Error validating body", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	product, err := pc.productService.CreateProduct(productDTO.Name, productDTO.Price, productDTO.CategoryID)
	if err != nil {
		logger.Error("Error saving product", err)
		restErr := rest_err.NewInternalServerError("Error saving product")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Product saved")
	return c.JSON(dtos.ProductToDTO(*product))
}

func (pc *ProductController) FindById(c *fiber.Ctx) error {
	logger.Info("Getting product by id")
	id := c.Params("id")

	if id == "" {
		logger.Error("product id is required", nil)
		restErr := rest_err.NewBadRequestError("product id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("product id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("product id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	product, err := pc.productService.GetProductByID(id)
	if err != nil {
		logger.Error("Error trying to get product by id", err)
		restErr := rest_err.NewInternalServerError("Error trying to get product by id")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Product found")
	return c.JSON(dtos.ProductToDTO(*product))
}

func (pc *ProductController) GetAll(c *fiber.Ctx) error {
	logger.Info("Getting all products")
	products, err := pc.productService.GetAllProducts()
	if err != nil {
		logger.Error("Error trying to get all products", err)
		restErr := rest_err.NewInternalServerError("Error trying to get all products")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Products found")
	return c.JSON(products)
}

func (pc *ProductController) Update(c *fiber.Ctx) error {
	logger.Info("Updating product")
	id := c.Params("id")
	var productDTO dtos.ProductDTO
	if err := c.BodyParser(&productDTO); err != nil {
		logger.Error("Error parsing body", err)
		restErr := rest_err.NewBadRequestError("invalid json body")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(productDTO); err != nil {
		logger.Error("Error validating body", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	product, err := pc.productService.UpdateProduct(id, productDTO.Name, productDTO.Price, productDTO.CategoryID)
	if err != nil {
		logger.Error("Error trying to update product", err)
		restErr := rest_err.NewInternalServerError("Error trying to update product")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Product updated")
	return c.JSON(dtos.ProductToDTO(*product))
}

func (pc *ProductController) Delete(c *fiber.Ctx) error {
	logger.Info("Deleting product")
	id := c.Params("id")

	if id == "" {
		logger.Error("product id is required", errors.New("product id is required"))
		restErr := rest_err.NewBadRequestError("product id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("product id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("product id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	err = pc.productService.DeleteProduct(id)
	if err != nil {
		logger.Error("Error trying to delete product", err)
		restErr := rest_err.NewInternalServerError("Error trying to delete product")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Product deleted")
	return c.SendStatus(http.StatusNoContent)
}
