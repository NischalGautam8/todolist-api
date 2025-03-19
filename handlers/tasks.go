package handlers

import (
	"go_todolist/db"
	"go_todolist/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	var tasks []models.Task
	rows, err := db.DB.Query("SELECT * FROM tasks")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Desce, &task.Completed)
		if err != nil {
			panic(err.Error())
		}
		tasks = append(tasks, task)
	}
	c.JSON(http.StatusOK, tasks)
}
