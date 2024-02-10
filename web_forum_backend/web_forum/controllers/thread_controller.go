// controllers/thread_controller.go
package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"web_forum_db/web_forum/models"

	"github.com/gin-gonic/gin"
)

type ThreadController struct {
	db *sql.DB
}

func NewThreadController(db *sql.DB) *ThreadController {
	return &ThreadController{db}
}

func (tc *ThreadController) GetThreadsByUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching threads"})
		return
	}
	threads, err := models.GetThreadsByUserIDFromDB(tc.db, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching threads"})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (tc *ThreadController) GetThreadsByTag(c *gin.Context) {
	tagName := c.Param("name")
	tag, err := models.GetTagByTagNameFromDB(tc.db, tagName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching tag"})
		return
	}

	if tag == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	threads, err := models.GetThreadsByTagIDFromDB(tc.db, tag.TagID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching threads"})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (tc *ThreadController) GetAllThreads(c *gin.Context) {
	threads, err := models.GetAllThreadsFromDB(tc.db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching threads"})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (tc *ThreadController) CreateThread(c *gin.Context) {
	var newThread models.Thread
	if err := c.ShouldBindJSON(&newThread); err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format, " + err.Error()})
		return
	}

	threadID, err := models.InsertThreadIntoDB(tc.db, &newThread)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding thread"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"thread_id": threadID})
}

func (tc *ThreadController) IncrementUpvoteCount(c *gin.Context) {
	threadIDParam := c.Param("id")
	threadID, err := strconv.Atoi(threadIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	err = models.IncrementDbUpvoteCount(tc.db, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error incrementing upvote count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Upvote count for thread %d incremented", threadID)})
}

func (tc *ThreadController) DecrementUpvoteCount(c *gin.Context) {
	threadIDParam := c.Param("id")
	threadID, err := strconv.Atoi(threadIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	err = models.DecrementDbUpvoteCount(tc.db, threadID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decrementing upvote count"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Upvote count for thread %d decreamented", threadID)})
}

func (tc *ThreadController) UpdateThreadContent(c *gin.Context) {
	threadIDParam := c.Param("id")
	content := c.Param("content")
	threadID, err := strconv.Atoi(threadIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	err = models.EditContentInDb(tc.db, threadID, content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating content"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Content edited for thread %d", threadID)})
}
