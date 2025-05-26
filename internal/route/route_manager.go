package route

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"empty/internal/delivery/http"
	"empty/internal/repository"
	"empty/internal/service"
)

func Init(r *gin.Engine, db *pgxpool.Pool) {
	userRoute := r.Group("/users")
	SetupUserRoutes(userRoute, db)
}

func SetupUserRoutes(r *gin.RouterGroup, db *pgxpool.Pool) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)
	{
		r.POST("/", userHandler.CreateUser)
		r.GET("/users/:id", userHandler.GetUserByID)
	}
}
