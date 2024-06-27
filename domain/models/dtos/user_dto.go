package dtos

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type UserDTO struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
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
	return models.NewUser(user.ID, user.Name, user.Email, user.Password)
}
