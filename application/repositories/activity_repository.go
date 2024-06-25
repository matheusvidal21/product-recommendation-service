package repositories

import "github.com/matheusvidal21/product-recommendation-service/domain/models"

type ActivityRepositoryInterface interface {
	SaveActivity(activity *models.UserActivityDoomain) (*models.UserActivityDoomain, error)
	GetActivityByUserId(userId string) ([]models.UserActivityDoomain, error)
}
