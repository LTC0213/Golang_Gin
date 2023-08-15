package controllers

import (
	"Golang_Gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskController struct {
	taskDAO *models.TaskDAO
}

func NewTaskController(taskDAO *models.TaskDAO) *TaskController {
	return &TaskController{taskDAO: taskDAO}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task
	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.taskDAO.Create(&task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) UpdateTask(ctx *gin.Context) {
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := c.taskDAO.FindByID(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = c.taskDAO.Update(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) GetTask(ctx *gin.Context) {
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := c.taskDAO.FindByID(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) GetAllTasks(ctx *gin.Context) {
	tasks, err := c.taskDAO.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) DeleteTask(ctx *gin.Context) {
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	task, err := c.taskDAO.FindByID(uint(taskID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	err = c.taskDAO.Delete(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
