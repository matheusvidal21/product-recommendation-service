package repositories

import (
	"context"
	"database/sql"
	"github.com/matheusvidal21/product-recommendation-service/domain/models"
	"github.com/matheusvidal21/product-recommendation-service/framework/database"
	logger "github.com/matheusvidal21/product-recommendation-service/framework/logging"
)

type UserRepositoryInterface interface {
	FindAll() ([]models.UserDomain, error)
	FindByID(id string) (*models.UserDomain, error)
	Create(user models.UserDomain) (*models.UserDomain, error)
	Update(id string, user models.UserDomain) (*models.UserDomain, error)
	Delete(id string) error
}

type UserRepository struct {
	queries *database.Queries
	ctx     context.Context
}

func NewUserRepository(db *sql.DB, ctx context.Context) UserRepositoryInterface {
	return &UserRepository{
		queries: database.New(db),
		ctx:     ctx,
	}
}

func (r *UserRepository) FindAll() ([]models.UserDomain, error) {
	logger.Info("Fetching all users")
	users, err := r.queries.GetAllUsers(r.ctx)
	if err != nil {
		logger.Error("Error fetching users", err)
		return nil, err
	}

	var userDomains []models.UserDomain
	for _, user := range users {
		userDomain := models.NewUserDomain(user.ID, user.Name, user.Email, "")
		userDomains = append(userDomains, userDomain)
	}

	logger.Info("Users fetched")
	return userDomains, nil
}

func (r *UserRepository) FindByID(id string) (*models.UserDomain, error) {
	logger.Info("Fetching user by ID")
	user, err := r.queries.GetUserByID(r.ctx, id)
	if err != nil {
		logger.Error("Error fetching user", err)
		return nil, err
	}

	userDomain := models.NewUserDomain(user.ID, user.Name, user.Email, "")
	return &userDomain, nil
}

func (r *UserRepository) Create(user models.UserDomain) (*models.UserDomain, error) {
	logger.Info("Creating user")
	pass, err := user.EncryptPassword()
	if err != nil {
		logger.Error("Error encrypting password", err)
		return nil, err
	}

	newUser, err := r.queries.CreateUser(r.ctx, database.CreateUserParams{
		ID:       user.GetID(),
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: pass,
	})

	if err != nil {
		logger.Error("Error creating user", err)
		return nil, err
	}

	logger.Info("User created")
	user = models.NewUserDomain(newUser.ID, newUser.Name, newUser.Email, "")
	return &user, nil
}

func (r *UserRepository) Update(id string, user models.UserDomain) (*models.UserDomain, error) {
	logger.Info("Updating user")
	pass, err := user.EncryptPassword()
	if err != nil {
		logger.Error("Error encrypting password", err)
		return nil, err
	}

	updatedUser, err := r.queries.UpdateUser(r.ctx, database.UpdateUserParams{
		ID:       id,
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Password: pass,
	})

	if err != nil {
		logger.Error("Error updating user", err)
		return nil, err
	}

	user = models.NewUserDomain(updatedUser.ID, updatedUser.Name, updatedUser.Email, "")
	logger.Info("User updated")
	return &user, nil
}

func (r *UserRepository) Delete(id string) error {
	logger.Info("Deleting user")
	err := r.queries.DeleteUser(r.ctx, id)
	if err != nil {
		logger.Error("Error deleting user", err)
		return err
	}

	logger.Info("User deleted")
	return nil
}
