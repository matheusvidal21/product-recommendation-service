package server

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gofiber/fiber/v2"
	recover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/matheusvidal21/product-recommendation-service/application/handlers"
	"github.com/matheusvidal21/product-recommendation-service/application/repositories"
	"github.com/matheusvidal21/product-recommendation-service/application/services"
	"github.com/matheusvidal21/product-recommendation-service/framework/config"
	"github.com/matheusvidal21/product-recommendation-service/framework/database"
)

type Server struct {
	App                *fiber.App
	ctx                context.Context
	Config             *config.Config
	Db                 *sql.DB
	CategoryController handlers.CategoryControllerInterface
	UserController     handlers.UserControllerInterface
	ProductController  handlers.ProductControllerInterface
	ActivityController handlers.ActivityControllerInterface
}

func NewServer() (*Server, error) {
	conf := config.LoadConfig()
	if conf == nil {
		return nil, errors.New("error loading config")
	}

	db, err := database.NewPostgresConnection(conf.PostgresURL, conf.DBPort, conf.DBUser, conf.DBPassword, conf.DBName)
	if err != nil {
		return nil, err
	}

	server := &Server{
		Db:     db,
		ctx:    context.Background(),
		Config: conf,
	}
	server.initDependecies()

	app := fiber.New()
	app.Use(recover.New())
	server.App = app

	server.initRoutes()
	return server, nil
}

func (s *Server) initRoutes() {
	userGroup := s.App.Group("/api/v1/users")
	userGroup.Get("/", s.UserController.FindAll)
	userGroup.Get("/:id", s.UserController.FindByID)
	userGroup.Post("/", s.UserController.Create)
	userGroup.Put("/:id", s.UserController.Update)
	userGroup.Delete("/:id", s.UserController.Delete)

	categoryGroup := s.App.Group("/api/v1/categories")
	categoryGroup.Get("/", s.CategoryController.GetAll)
	categoryGroup.Get("/:id", s.CategoryController.FindById)
	categoryGroup.Post("/", s.CategoryController.Save)
	categoryGroup.Put("/:id", s.CategoryController.Update)
	categoryGroup.Delete("/:id", s.CategoryController.Delete)

	productGroup := s.App.Group("/api/v1/products")
	productGroup.Get("/", s.ProductController.GetAll)
	productGroup.Get("/:id", s.ProductController.FindById)
	productGroup.Post("/", s.ProductController.Save)
	productGroup.Put("/:id", s.ProductController.Update)
	productGroup.Delete("/:id", s.ProductController.Delete)

	activityGroup := s.App.Group("/api/v1/activities")
	activityGroup.Get("/", s.ActivityController.GetAll)
	activityGroup.Get("/:userId", s.ActivityController.GetActivityByUserId)
	activityGroup.Post("/", s.ActivityController.Save)
}

func (s *Server) initDependecies() {
	userRepo := repositories.NewUserRepository(s.Db, s.ctx)
	userService := services.NewUserService(userRepo, s.ctx)
	s.UserController = handlers.NewUserController(userService)

	categoryRepo := repositories.NewCategoryRepository(s.Db, s.ctx)
	categoryService := services.NewCategoryService(categoryRepo, s.ctx)
	s.CategoryController = handlers.NewCategoryController(categoryService)

	productRepo := repositories.NewProductRepository(s.Db, s.ctx)
	productService := services.NewProductService(productRepo, categoryService, s.ctx)
	s.ProductController = handlers.NewProductController(productService)

	activityRepo := repositories.NewActivityRepository(s.Db, s.ctx)
	activityService := services.NewActivityService(activityRepo, s.ctx, productService, userService)
	s.ActivityController = handlers.NewActivityController(activityService)
}

func (s *Server) Run() error {
	return s.App.Listen(s.Config.AppPort)
}
