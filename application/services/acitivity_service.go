package services

import (
	"context"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/domain/models/dtos"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
)

type ActivityServiceInterface interface {
	SaveActivity(userID, productID, action string) (*models.UserActivityDomain, error)
	GetActivityByUserId(userID string) ([]dtos.UserActivityDTO, error)
	GetAllActivities() ([]dtos.UserActivityDTO, error)
}

type ActivityService struct {
	userService    UserServiceInterface
	productService ProductServiceInterface
	repo           repositories.ActivityRepositoryInterface
	ctx            context.Context
}

func NewActivityService(repo repositories.ActivityRepositoryInterface, ctx context.Context, productService ProductServiceInterface, userService UserServiceInterface) ActivityServiceInterface {
	return &ActivityService{
		userService:    userService,
		productService: productService,
		repo:           repo,
		ctx:            ctx,
	}
}

func (s *ActivityService) SaveActivity(userID, productID, action string) (*models.UserActivityDomain, error) {
	logger.Info("Saving activity services")

	user, err := s.userService.GetUserByID(userID)
	if err != nil || user == nil {
		logger.Error("Error getting user by id", err)
		return nil, err
	}

	product, err := s.productService.GetProductByID(productID)
	if err != nil || product == nil {
		logger.Error("Error getting product by id", err)
		return nil, err
	}

	activity := models.NewUserActivity(userID, productID, models.StringParseAction(action))
	return s.repo.SaveActivity(activity)
}

func (s *ActivityService) GetActivityByUserId(userID string) ([]dtos.UserActivityDTO, error) {
	logger.Info("Getting activity by user id services")
	act := []dtos.UserActivityDTO{}

	acitivities, err := s.repo.GetActivityByUserId(userID)
	if err != nil {
		return act, err
	}

	for _, a := range acitivities {
		act = append(act, dtos.UserActivityToDTO(a))
	}

	return act, nil
}

func (s *ActivityService) GetAllActivities() ([]dtos.UserActivityDTO, error) {
	logger.Info("Getting all activities services")
	act := []dtos.UserActivityDTO{}

	acitivities, err := s.repo.FindAll()
	if err != nil {
		return act, err
	}

	for _, a := range acitivities {
		act = append(act, dtos.UserActivityToDTO(a))
	}

	return act, nil
}
