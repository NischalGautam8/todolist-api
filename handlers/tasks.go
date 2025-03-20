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

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.DB.Exec("INSERT INTO task (title, desce, completed) VALUES ($1, $2, $3)", task.Title, task.Desce, task.Completed)
	if err != nil {
		panic(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}
func UpdateTask(c *gin.Context) {
	var task models.Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := db.DB.Exec("UPDATE task SET title = $1, desce = $2, completed = $3 WHERE id = $4", task.Title, task.Desce, task.Completed, task.Id)
	if err != nil {
		panic(err.Error())
	}

}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("delete from task where id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
