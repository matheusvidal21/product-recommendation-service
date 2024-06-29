package dtos

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type UserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6,max=100"`
}

func UserToDTO(user models.UserDomain) UserDTO {
	return UserDTO{
		ID:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: user.GetPassword(),
	}
}

func UserToDomain(user UserDTO) models.UserDomain {
	return models.NewUserWithId(user.ID, user.Name, user.Email, user.Password)
}

type UserGetAllDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserCreateDTO struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name" validate:"required,min=3,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password,omitempty" validate:"required,min=6,max=100"`
}
