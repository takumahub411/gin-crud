package main

import (
	"gin-curd/controllers"
	"gin-curd/initializers"
	"gin-curd/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initializers.LoadEnvfile()
	initializers.ConnectDb()
	initializers.SyncDatabase()

}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	r.GET("/validate", middleware.GetAuthFromCookie, controllers.Validate)

	r.POST("/movies", controllers.AddMovie)
	r.GET("/movies", controllers.GetMovies)
	r.GET("/movies/:id", controllers.GetMovie)
	r.PUT("/movies/:id", controllers.UpdateMovie)
	r.DELETE("/movies/:id", controllers.DeleteMovie)

	r.Run()
}
