// controllers/tag_controller.go
package controllers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	"web_forum_db/web_forum/models"
)

type TagController struct {
	db *sql.DB
}

func NewTagController(db *sql.DB) *TagController {
	return &TagController{db}
}

func (tagC *TagController) GetAllTag(c *gin.Context) {
	threads, err := models.GetAllThreadsFromDB(tagC.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tags"})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (tc *TagController) TagThread(c *gin.Context) {
	var newThreadTag models.ThreadTag
	if err := c.ShouldBindJSON(&newThreadTag); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format, " + err.Error()})
		return
	}

	err := models.InsertThreadTagIntoDB(tc.db, &newThreadTag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error tagging thread"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})
}
