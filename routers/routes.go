package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/config"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/controllers"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/middlewares"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	config.ConnectToDB()

	// Auth
	router.GET("/auth/login", controllers.LoginUser)
	router.POST("/auth/register", controllers.RegisterUser)
	router.GET("/auth/logout", controllers.LogoutUser)

	// User
	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserById)
	router.PUT("/users", middlewares.RequireAuth, controllers.UpdateUser)
	router.DELETE("/users", middlewares.RequireAuth, controllers.DeleteUser)

	// Photo
	router.GET("/photos", controllers.GetPhotos)
	router.GET("/photos/:id", controllers.GetPhotoById)
	router.POST("/photos", middlewares.RequireAuth, controllers.CreatePhoto)
	router.PUT("/photos/:id", middlewares.RequireAuth, controllers.UpdatePhoto)
	router.DELETE("/photos/:id", middlewares.RequireAuth, controllers.DeletePhoto)

	return router
}
