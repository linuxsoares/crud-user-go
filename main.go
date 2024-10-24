package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/linuxsoares/crud-user-go/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "John"},
	{ID: 2, Name: "Doe"},
}

func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Users API"
	docs.SwaggerInfo.Description = "User API Server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r.GET("/users", getUsers)

	r.GET("/users/:id", getUser)

	r.POST("/users", createUser)

	r.PUT("/users", editUser)

	r.DELETE("/users/:id", removeUser)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}

// @BasePath /
// ShowAccount 	 godoc
// @Summary      Show a Users list
// @Description  get users
// @Tags         getUsers
// @Accept       json
// @Produce      json
// @Router       /users  [get]
func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// @BasePath /
// ShowAccount 	 godoc
// @Summary      Show a Users list
// @Description  get users
// @Tags         getUsers
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "ID"
// @Router       /users/{id}  [get]
func getUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	userFound, _ := findUserByID(users, id)
	c.JSON(http.StatusOK, userFound)
}

// @BasePath /
// ShowAccount 	 godoc
// @Summary      Create a User
// @Description  create users
// @Tags         createUser
// @Accept       json
// @Produce      json
// @Param user body User true "User to create"
// @Router 		/users [post]
func createUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	users = append(users, user)
	c.JSON(http.StatusCreated, user)
}

// @BasePath /
// ShowAccount 	 godoc
// @Summary      Edit a User
// @Description  edit users
// @Tags         editUser
// @Accept       json
// @Produce      json
// @Param user body User true "User to edit"
// @Router 		/users [put]
func editUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	putUser(users, user)
	c.JSON(http.StatusOK, user)
}

// @BasePath /
// ShowAccount 	 godoc
// @Summary      Remove a Users list
// @Description  delete user
// @Tags         removeUser
// @Accept       json
// @Produce      json
// @Param        id         path      string  true  "ID"
// @Router       /users/{id}  [delete]
func removeUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	deleteUserByID(users, id)
	c.JSON(http.StatusOK, nil)
}

func findUserByID(users []User, id int) (*User, bool) {
	for _, user := range users {
		if user.ID == id {
			return &user, true
		}
	}
	return nil, false
}

func deleteUserByID(users []User, id int) []User {
	newUsers := []User{}
	for _, user := range users {
		if user.ID != id {
			newUsers = append(newUsers, user)
		}
	}
	return newUsers
}

func putUser(users []User, newUser User) []User {
	for i := range users {
		if users[i].ID == newUser.ID {
			// Update the existing user
			users[i] = newUser
			return users
		}
	}
	// Add new user if not found
	return append(users, newUser)
}
