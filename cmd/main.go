package main

import (
	"github.com/aafak/auth-service/internal/handler"
	"github.com/aafak/auth-service/internal/repository"
	"github.com/aafak/auth-service/internal/service"
	"github.com/gin-gonic/gin"
)

// type User struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// }

// var users = []User{
// 	{ID: "1", Name: "user1"},
// 	{ID: "2", Name: "user2"},
// 	{ID: "3", Name: "user3"},
// }

// func getUsers(c *gin.Context) {
// 	name := c.Query("name") // query parameter
// 	if name != "" {
// 		for _, user := range users {
// 			if user.Name == name {
// 				c.IndentedJSON(http.StatusOK, user)
// 				return
// 			}
// 		}
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
// 		return
// 	} else {
// 		//c.JSON(200, books)
// 		c.IndentedJSON(http.StatusOK, users)
// 	}
// }

func main() {
	db, err := repository.NewPostgresDB("")
	if err != nil {
		panic("Failed to connect to DB")
	}
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	router.POST("/register", userHandler.RegisterUser)
	//POST  http://localhost:8080/register      { "username" : "aafak",  "password": "test"}
	// http://localhost:8080/users?name=user1
	router.Run(":8080")
}
