package services

import (
	"context"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
)

type ActivityServiceInterface interface {
	SaveActivity(userID, productID, action string) (*models.UserActivityDomain, error)
	GetActivityByUserId(userID string) ([]models.UserActivityDomain, error)
	GetAllActivities() ([]models.UserActivityDomain, error)
}

type ActivityService struct {
	repo repositories.ActivityRepositoryInterface
	ctx  context.Context
}

func NewActivityService(repo repositories.ActivityRepositoryInterface, ctx context.Context) ActivityServiceInterface {
	return &ActivityService{
		repo: repo,
		ctx:  ctx,
	}
}

func (s *ActivityService) SaveActivity(userID, productID, action string) (*models.UserActivityDomain, error) {
	activity := models.NewUserActivity(userID, productID, models.StringParseAction(action))
	return s.repo.SaveActivity(activity)
}

func (s *ActivityService) GetActivityByUserId(userID string) ([]models.UserActivityDomain, error) {
	return s.repo.GetActivityByUserId(userID)
}

func (s *ActivityService) GetAllActivities() ([]models.UserActivityDomain, error) {
	return s.repo.FindAll()
}
