package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/matheusvidal21/product-recommendation-service/application/services"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
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
	users, err := uc.userService.GetAllUsers()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(users)
}

func (uc *UserController) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	user, err := uc.userService.GetUserByID(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var userDTO dtos.UserCreateDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(userDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uc.userService.CreateUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(dtos.UserToDTO(*user))
}

func (uc *UserController) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	var userDTO dtos.UserDTO
	if err := c.BodyParser(&userDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := validate.Struct(userDTO); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := uc.userService.UpdateUser(id, userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (uc *UserController) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is required"})
	}

	_, err := uuid.Parse(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "id is not a valid UUID"})
	}

	if err := uc.userService.DeleteUser(id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusNoContent)
}
