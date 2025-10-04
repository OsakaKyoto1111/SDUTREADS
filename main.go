package main

import (
	"log"
	"myapp/internal/controller"
	"myapp/internal/middleware"
	"myapp/internal/model"
	"myapp/internal/repository"
	"myapp/internal/service"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=myuser password=mypassword dbname=myapp port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ DB connection failed:", err)
	}

	db.AutoMigrate(&model.User{}, &model.Post{})

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, "secret123")
	authHandler := controller.NewAuthHandler(authService)
	userHandler := controller.NewUserHandler(userRepo)

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := controller.NewPostHandler(postService)

	likeRepo := repository.NewLikeRepository(db)
	likeService := service.NewLikeService(likeRepo, postRepo)
	likeHandler := controller.NewLikeHandler(likeService)

	followRepo := repository.NewFollowRepository(db)
	followService := service.NewFollowService(followRepo, postRepo, userRepo)
	followHandler := controller.NewFollowHandler(followService)

	e := echo.New()

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	e.GET("/users", authHandler.GetAll)
	e.GET("/posts", postHandler.GetAll)
	e.GET("/posts/:id", postHandler.GetByID)

	auth := e.Group("")
	auth.Use(middleware.JWTMiddleware)
	auth.GET("/me", userHandler.Me)
	auth.POST("/posts", postHandler.Create)
	auth.DELETE("/posts/:id", postHandler.Delete)

	auth.POST("/posts/:id/like", likeHandler.Toggle)
	auth.DELETE("/posts/:id/like", likeHandler.Toggle)
	auth.GET("/posts/:id/likes", likeHandler.Count)

	auth.POST("/users/:id/follow", followHandler.ToggleFollow)
	auth.DELETE("/users/:id/follow", followHandler.ToggleFollow)
	auth.GET("/feed", followHandler.Feed)

	log.Println("🚀 Server started at :8080")
	e.Start(":8080")
}
