// controllers/user_controller.go
package controllers

import (
	"database/sql"
	"net/http"
	"web_forum_db/web_forum/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{db}
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user, err := models.GetUserByUsernameFromDB(uc.db, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
		println(err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format, " + err.Error()})
        return
    }

    userID, err := models.InsertUserIntoDB(uc.db, &newUser)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding user"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"user_id": userID})
}
