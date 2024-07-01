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

type UserControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type UserController struct {
	userService services.UserServiceInterface
}

func NewUserController(userService services.UserServiceInterface) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) FindAll(c *fiber.Ctx) error {
	logger.Info("FindAll users")
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		logger.Error("Error when trying to get users", err)
		restErr := rest_err.NewInternalServerError("Error when trying to get users")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("Users found")
	return c.JSON(users)
}

func (uc *UserController) FindByID(c *fiber.Ctx) error {
	logger.Info("FindByID user")
	id := c.Params("id")

	if id == "" {
		logger.Error("user id is required", errors.New("user id is required"))
		restErr := rest_err.NewBadRequestError("user id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("user id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("user id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		logger.Error("Error when trying to get user", err)
		restErr := rest_err.NewInternalServerError("Error when trying to get user")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("User found")
	return c.JSON(dtos.UserToResponseDTO(*user))
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	logger.Info("Create user")
	var userDTO dtos.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		logger.Error("Error when trying to parse user", err)
		restErr := rest_err.NewBadRequestError("Error when trying to parse user")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(userDTO); err != nil {
		logger.Error("Error when trying to validate user", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	user, err := uc.userService.CreateUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		logger.Error("Error when trying to create user", err)
		restErr := rest_err.NewInternalServerError("Error when trying to create user")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("User created")
	return c.JSON(dtos.UserToResponseDTO(*user))
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	logger.Info("Update user")
	id := c.Params("id")
	if id == "" {
		logger.Error("user id is required", errors.New("user id is required"))
		restErr := rest_err.NewBadRequestError("user id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("user id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("user id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	var userDTO dtos.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		logger.Error("Error when trying to parse user", err)
		restErr := rest_err.NewBadRequestError("Error when trying to parse user")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := validation.Validate.Struct(userDTO); err != nil {
		logger.Error("Error when trying to validate user", err)
		restErr := validation.ValidateStructError(err)
		return c.Status(restErr.Code).JSON(restErr)
	}

	user, err := uc.userService.UpdateUser(id, userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		logger.Error("Error when trying to update user", err)
		restErr := rest_err.NewInternalServerError("Error when trying to update user")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("User updated")
	return c.JSON(dtos.UserToResponseDTO(*user))
}

func (uc *UserController) Delete(c *fiber.Ctx) error {
	logger.Info("Delete user")
	id := c.Params("id")
	if id == "" {
		logger.Error("user id is required", errors.New("user id is required"))
		restErr := rest_err.NewBadRequestError("user id is required")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		logger.Error("id is not a valid UUID", err)
		restErr := rest_err.NewBadRequestError("id is not a valid UUID")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	if err := uc.userService.DeleteUser(id); err != nil {
		logger.Error("Error when trying to delete user", err)
		restErr := rest_err.NewInternalServerError("Error when trying to delete user")
		return c.Status(restErr.Code).JSON(fiber.Map{"error": restErr.Message})
	}

	logger.Info("User deleted")
	return c.SendStatus(http.StatusNoContent)
}
