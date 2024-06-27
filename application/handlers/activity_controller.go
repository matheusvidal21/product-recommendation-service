package handlers

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/services"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	"net/http"
)

var validate = validator.New()

type ActivityControllerInterface interface {
	Save(c *fiber.Ctx) error
	GetActivityByUserId(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

type ActivityController struct {
	activityService services.ActivityServiceInterface
}

func NewActivityController(activityService services.ActivityServiceInterface) *ActivityController {
	return &ActivityController{activityService: activityService}
}

func (ac *ActivityController) Save(c *fiber.Ctx) error {
	var activityDTO dtos.UserActivityDTO
	if err := c.BodyParser(&activityDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(activityDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	activity, err := ac.activityService.SaveActivity(activityDTO.UserID, activityDTO.ProductID, activityDTO.Action)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(activity)
}

func (ac *ActivityController) GetActivityByUserId(c *fiber.Ctx) error {
	id := c.Params("userId")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "userId is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "userId is not a valid UUID"})
	}

	activities, err := ac.activityService.GetActivityByUserId(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(activities)
}

func (ac *ActivityController) GetAll(c *fiber.Ctx) error {
	activities, err := ac.activityService.GetAllActivities()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(activities)
}
