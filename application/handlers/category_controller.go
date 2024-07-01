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
	logger.Info("Saving category")
	var categoryDTO dtos.CategoryDTO
	if err := c.BodyParser(&categoryDTO); err != nil {
		logger.Error("Error trying to parse body", err)
		restErr := rest_err.NewBadRequestError("Invalid request body")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(categoryDTO); err != nil {
		logger.Error("Error trying to validate struct", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	category, err := cc.categoryService.CreateCategory(categoryDTO.Name, categoryDTO.Description)
	if err != nil {
		logger.Error("Error trying to save category", err)
		restErr := rest_err.NewInternalServerError("Error trying to save category")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Category saved")
	return c.JSON(dtos.CategoryToDTO(*category))
}

func (cc *CategoryController) FindById(c *fiber.Ctx) error {
	logger.Info("Getting category by id")
	id := c.Params("id")

	if id == "" {
		logger.Error("category id is required", errors.New("category id is required"))
		restErr := rest_err.NewBadRequestError("category id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("category id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("category id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	category, err := cc.categoryService.GetCategoryByID(id)
	if err != nil {
		logger.Error("Error trying to get category by id", err)
		restErr := rest_err.NewInternalServerError("Error trying to get category by id")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Category found")
	return c.JSON(dtos.CategoryToDTO(*category))
}

func (cc *CategoryController) GetAll(c *fiber.Ctx) error {
	logger.Info("Getting all categories")
	categories, err := cc.categoryService.GetAllCategories()
	if err != nil {
		logger.Error("Error trying to get all categories", err)
		restErr := rest_err.NewInternalServerError("Error trying to get all categories")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Categories found")
	return c.JSON(categories)
}

func (cc *CategoryController) Update(c *fiber.Ctx) error {
	logger.Info("Updating category")
	id := c.Params("id")
	var categoryDTO dtos.CategoryDTO
	if err := c.BodyParser(&categoryDTO); err != nil {
		logger.Error("Error trying to parse body", err)
		restErr := rest_err.NewBadRequestError("Invalid request body")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(categoryDTO); err != nil {
		logger.Error("Error trying to validate struct", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	category, err := cc.categoryService.UpdateCategory(id, categoryDTO.Name, categoryDTO.Description)
	if err != nil {
		logger.Error("Error trying to update category", err)
		restErr := rest_err.NewInternalServerError("Error trying to update category")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Category updated")
	return c.JSON(dtos.CategoryToDTO(*category))
}

func (cc *CategoryController) Delete(c *fiber.Ctx) error {
	logger.Info("Deleting category")
	id := c.Params("id")

	if id == "" {
		logger.Error("id is required", errors.New("category id is required"))
		restErr := rest_err.NewBadRequestError("category id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("category id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("category id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	err = cc.categoryService.DeleteCategory(id)
	if err != nil {
		logger.Error("Error trying to delete category", err)
		restErr := rest_err.NewInternalServerError("Error trying to delete category")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Category deleted")
	return c.SendStatus(http.StatusNoContent)
}
