package repositories

import (
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"strings"
)

var USER_ACTIVITY_INDEX = "user_activity"

type ActivityRepositoryInterface interface {
	SaveActivity(activity models.UserActivityDomain) (*models.UserActivityDomain, error)
	GetActivityByUserId(userId string) ([]models.UserActivityDomain, error)
	FindAll() ([]models.UserActivityDomain, error)
}

type ActivityRepository struct {
	client *elasticsearch.Client
	ctx    context.Context
}

func NewActivityRepository(client *elasticsearch.Client, ctx context.Context) *ActivityRepository {
	return &ActivityRepository{
		client: client,
		ctx:    ctx,
	}
}

func (r *ActivityRepository) SaveActivity(activity models.UserActivityDomain) (*models.UserActivityDomain, error) {
	activity = models.NewUserActivity(activity.GetUserID(), activity.GetProductID(), models.StringParseAction(activity.GetAction()))
	body, err := json.Marshal(activity)
	if err != nil {
		return nil, err
	}

	res, err := r.client.Index(
		USER_ACTIVITY_INDEX,
		strings.NewReader(string(body)),
		r.client.Index.WithDocumentID(activity.GetUserID()+"_"+activity.GetProductID()),
		r.client.Index.WithContext(r.ctx),
	)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, err
	}

	return &activity, nil
}

func (r *ActivityRepository) GetActivityByUserId(userId string) ([]models.UserActivityDomain, error) {
	query := `{"query": {"match": {"user_id": "` + userId + `"}}}`
	res, err := r.client.Search(
		r.client.Search.WithContext(r.ctx),
		r.client.Search.WithIndex(USER_ACTIVITY_INDEX),
		r.client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		return nil, err
	}

	var activities []models.UserActivityDomain
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		var activity models.UserActivityDomain
		hitSource := hit.(map[string]interface{})["_source"]
		source, err := json.Marshal(hitSource)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(source, &activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func (r *ActivityRepository) FindAll() ([]models.UserActivityDomain, error) {
	res, err := r.client.Search(
		r.client.Search.WithContext(r.ctx),
		r.client.Search.WithIndex(USER_ACTIVITY_INDEX),
	)
	if err != nil {
		return nil, err
	}

	var activities []models.UserActivityDomain
	var result map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
	for _, hit := range hits {
		var activity models.UserActivityDomain
		hitSource := hit.(map[string]interface{})["_source"]
		source, err := json.Marshal(hitSource)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(source, &activity)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}
