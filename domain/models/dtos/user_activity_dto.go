package dtos

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type UserActivityDTO struct {
	UserID    string `json:"user_id" validate:"required,uuid"`
	ProductID string `json:"product_id" validate:"required,uuid"`
	Action    string `json:"action" validate:"required,validAction"`
}

func UserActivityToDTO(userActivity models.UserActivityDomain) UserActivityDTO {
	return UserActivityDTO{
		UserID:    userActivity.GetUserID(),
		ProductID: userActivity.GetProductID(),
		Action:    userActivity.GetAction(),
	}
}

func UserActivityToDomain(userActivity UserActivityDTO) models.UserActivityDomain {
	return models.NewUserActivity(userActivity.UserID, userActivity.ProductID, models.StringParseAction(userActivity.Action))
}
