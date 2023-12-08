package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Represents a task .

// Tasks Struct

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

// Mock data for tasks

var tasks = []Task{
	{ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
	{ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
	{ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// Update Task Func

func taskUpdate(ctx *gin.Context) {
	id := ctx.Param("id")
	var updatedTask Task

	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, task := range tasks {
		if task.ID == id {
			// Update only the specified fields
			if updatedTask.Title != "" {
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			ctx.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Get Task Get by ID

func taskId(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, task := range tasks {
		if task.ID == id {
			ctx.JSON(http.StatusOK, task)
			return
		}

	}
	ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})

}

// Delete Task Function

func taskDelete(ctx *gin.Context) {
	id := ctx.Param("id")

	for i, valu := range tasks {
		if valu.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})
			return
		}
	}
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Post New Task Function

func postTask(ctx *gin.Context) {
	var newTask Task
	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tasks = append(tasks, newTask)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}

// Get All Tasks function

func taskAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// Main Function

func main() {
	r := gin.Default()
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong Fathi",
		})
	})

	// Get All Task
	r.GET("/tasks", taskAll)

	// Route for TaskUpdate
	r.PUT("/tasks/:id", taskUpdate)

	// Route for Get Task By ID
	r.GET("/tasks/:id", taskId)

	// Route for Delete Task
	r.DELETE("/tasks/:id", taskDelete)

	// Route for PostTask
	r.POST("/tasks", postTask)

	// listen and serve on 0.0.0.0:8080
	r.Run()
}
