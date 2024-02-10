// main.go
package main

import (
	"database/sql"
	"fmt"
	"web_forum_db/web_forum/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/webforum")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Connected to the database")

	router := gin.Default()

	userController := controllers.NewUserController(db)
	threadController := controllers.NewThreadController(db)
	tagController := controllers.NewTagController(db)
	commentController := controllers.NewCommentController(db)

	router.GET("/user/:username", userController.GetUserByUsername)
	router.GET("/thread/user/:id", threadController.GetThreadsByUser)
	router.GET("/thread/tag/:name", threadController.GetThreadsByTag)
	router.GET("/tag", tagController.GetAllTag)
	router.GET("/thread", threadController.GetAllThreads)
	router.GET("/comment/:id/thread", commentController.GetCommentsByThread)

	router.POST("/thread", threadController.CreateThread)
	router.POST("/user", userController.CreateUser)
	router.POST("/tag", tagController.TagThread)
	router.POST("/comment", commentController.CreateComment)

	router.PUT("/thread/upvote/:id", threadController.IncrementUpvoteCount)
	router.PUT("/thread/downvote/:id", threadController.DecrementUpvoteCount)
	router.PUT("/thread/update/:id/:content", threadController.UpdateThreadContent)

	

	router.Run(":8080")
}
