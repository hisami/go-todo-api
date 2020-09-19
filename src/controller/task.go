package controller

import (
	"github.com/gin-gonic/gin"
	"../model"
	"time"
	"net/http"
	"fmt"
)

// タスク一覧
func TasksGET(c *gin.Context) {
	db := model.DBConnect()
	result, err := db.Query("SELECT * FROM task ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	tasks := []model.Task{}
	for result.Next() {
		task := model.Task{}
		var id uint
		var createdAt, updatedAt time.Time
		var title string

		err = result.Scan(&id, &createdAt, &updatedAt, &title)
		if err != nil {
			panic(err.Error())
		}

		task.ID = id
		task.CreatedAt = createdAt
		task.UpdatedAt = updatedAt
		task.Title = title
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// タスク登録
func TaskPOST(c *gin.Context) {
	db := model.DBConnect()

	title := c.PostForm("title")
	now := time.Now()

	_, err := db.Exec("INSERT INTO task(title,created_at,updated_at) VALUES(?,?,?)", title, now, now)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("post sent. title: %s", title)
}
