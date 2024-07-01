package repositories

import (
	"context"
	"database/sql"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/config/logging"
	"github.com/matheusvidal21/product-recommendation-service/framework/database"
)

type ActivityRepositoryInterface interface {
	SaveActivity(activity models.UserActivityDomain) (*models.UserActivityDomain, error)
	GetActivityByUserId(userId string) ([]models.UserActivityDomain, error)
	FindAll() ([]models.UserActivityDomain, error)
}

type ActivityRepository struct {
	queries *database.Queries
	ctx     context.Context
}

func NewActivityRepository(db *sql.DB, ctx context.Context) ActivityRepositoryInterface {
	return &ActivityRepository{
		queries: database.New(db),
		ctx:     ctx,
	}
}

func (r *ActivityRepository) SaveActivity(activity models.UserActivityDomain) (*models.UserActivityDomain, error) {
	logger.Info("Saving activity")
	err := r.queries.SaveActivity(r.ctx, database.SaveActivityParams{
		UserID:    activity.GetUserID(),
		ProductID: activity.GetProductID(),
		Action:    activity.GetAction(),
	})
	if err != nil {
		logger.Error("Error saving activity", err)
		return nil, err
	}

	logger.Info("Activity saved")
	return &activity, nil
}

func (r *ActivityRepository) GetActivityByUserId(userId string) ([]models.UserActivityDomain, error) {
	logger.Info("Fetching activities by user ID")
	activities, err := r.queries.GetActivityByUserId(r.ctx, userId)
	if err != nil {
		logger.Error("Error fetching activities", err)
		return nil, err
	}

	var activityDomains []models.UserActivityDomain
	for _, activity := range activities {
		activityDomain := models.NewUserActivity(activity.UserID, activity.ProductID, models.StringParseAction(activity.Action))
		activityDomains = append(activityDomains, activityDomain)
	}

	logger.Info("Activities fetched")
	return activityDomains, nil
}

func (r *ActivityRepository) FindAll() ([]models.UserActivityDomain, error) {
	logger.Info("Fetching all activities")
	activities, err := r.queries.GetAllActivities(r.ctx)
	if err != nil {
		logger.Error("Error fetching activities", err)
		return nil, err
	}

	var activityDomains []models.UserActivityDomain
	for _, activity := range activities {
		activityDomain := models.NewUserActivity(activity.UserID, activity.ProductID, models.StringParseAction(activity.Action))
		activityDomains = append(activityDomains, activityDomain)
	}

	logger.Info("Activities fetched")
	return activityDomains, nil
}
