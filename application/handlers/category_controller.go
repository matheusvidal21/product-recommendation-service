package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/services"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	"net/http"
)

type CategoryControllerInterface interface {
	Save(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type CategoryController struct {
	categoryService services.CategoryServiceInterface
}

func NewCategoryController(categoryService services.CategoryServiceInterface) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

func (cc *CategoryController) Save(c *fiber.Ctx) error {
	var categoryDTO dtos.CategoryDTO
	if err := c.BodyParser(&categoryDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(categoryDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	category, err := cc.categoryService.CreateCategory(categoryDTO.Name, categoryDTO.Description)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(category)
}

func (cc *CategoryController) FindById(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	category, err := cc.categoryService.GetCategoryByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(category)
}

func (cc *CategoryController) GetAll(c *fiber.Ctx) error {
	categories, err := cc.categoryService.GetAllCategories()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(categories)
}

func (cc *CategoryController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var categoryDTO dtos.CategoryDTO
	if err := c.BodyParser(&categoryDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(categoryDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	category, err := cc.categoryService.UpdateCategory(id, categoryDTO.Name, categoryDTO.Description)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(category)
}

func (cc *CategoryController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	err = cc.categoryService.DeleteCategory(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
