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
)

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
	logger.Info("Saving activity")
	var activityDTO dtos.UserActivityDTO

	if err := c.BodyParser(&activityDTO); err != nil {
		logger.Error("Error trying to parse body", err)
		restErr := rest_err.NewBadRequestError("Invalid request body")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(activityDTO); err != nil {
		logger.Error("Error trying to validate struct", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	activity, err := ac.activityService.SaveActivity(activityDTO.UserID, activityDTO.ProductID, activityDTO.Action)
	if err != nil {
		logger.Error("Error trying to save activity", err)
		restErr := rest_err.NewInternalServerError("Error trying to save activity")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Activity saved")
	return c.JSON(dtos.UserActivityToDTO(*activity))
}

func (ac *ActivityController) GetActivityByUserId(c *fiber.Ctx) error {
	logger.Info("Getting activities by user id")
	id := c.Params("userId")

	if id == "" {
		logger.Error("userId is required", errors.New("userId is required"))
		restErr := rest_err.NewBadRequestError("userId is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("userId is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("userId is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	activities, err := ac.activityService.GetActivityByUserId(id)
	if err != nil {
		logger.Error("Error trying to get activities by user id", err)
		restErr := rest_err.NewInternalServerError("Error trying to get activities by user id")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Activities found")
	return c.JSON(activities)
}

func (ac *ActivityController) GetAll(c *fiber.Ctx) error {
	logger.Info("Getting all activities")
	activities, err := ac.activityService.GetAllActivities()
	if err != nil {
		logger.Error("Error trying to get all activities", err)
		restErr := rest_err.NewInternalServerError("Error trying to get all activities")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Activities found")
	return c.JSON(activities)
}
