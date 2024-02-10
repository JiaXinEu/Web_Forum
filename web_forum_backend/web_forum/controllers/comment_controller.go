// controllers/comment_controller.go
package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"web_forum_db/web_forum/models"

	"github.com/gin-gonic/gin"
)

type CommentController struct {
	db *sql.DB
}

func NewCommentController(db *sql.DB) *CommentController {
	return &CommentController{db}
}

func (cc *CommentController) GetCommentsByThread(c *gin.Context) {
	threadID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching comments"})
		return
	}
	comments, err := models.GetCommentsByThreadIDFromDB(cc.db, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *CommentController) CreateComment(c *gin.Context) {
    var newComment models.Comment
    if err := c.ShouldBindJSON(&newComment); err != nil {
		println(err.Error())
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format, " + err.Error()})
        return
    }

    commentID, err := models.InsertCommentIntoDB(cc.db, &newComment)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding thread"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"comment_id": commentID})
}