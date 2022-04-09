package main

import (
	"net/http"

	api "goapi/api"
	nd "goapi/operations"

	"github.com/gin-gonic/gin"
)

func main() {
	nd.NepseDetails()
	r := setupRouter()
	_ = r.Run(":8080")

}
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	userRepo := api.New()
	// r.POST("/users", userRepo.CreateUser)
	r.GET("/stocks", userRepo.GetStocks)
	r.GET("/broker", userRepo.GetBroker)
	r.GET("/sector", userRepo.GetSector)
	// r.GET("/users/:id", userRepo.GetUser)
	// r.PUT("/users/:id", userRepo.UpdateUser)
	// r.DELETE("/users/:id", userRepo.DeleteUser)

	return r
}
