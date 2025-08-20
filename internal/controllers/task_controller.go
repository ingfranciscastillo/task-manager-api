package controllers

import (
	"net/http"
	"strconv"
	"task-manager/internal/models"
	"task-manager/internal/services"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	taskService *services.TaskService
}

func NewTaskController() *TaskController {
	return &TaskController{
		taskService: &services.TaskService{},
	}
}

// CreateTaskRequest representa la estructura de request para crear tarea
type CreateTaskRequest struct {
	Title       string            `json:"title" binding:"required"`
	Description string            `json:"description"`
	Status      models.TaskStatus `json:"status"`
}

// UpdateTaskRequest representa la estructura de request para actualizar tarea
type UpdateTaskRequest struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      models.TaskStatus `json:"status"`
}

// GetTasks obtiene todas las tareas del usuario autenticado
func (tc *TaskController) GetTasks(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	tasks, err := tc.taskService.GetUserTasks(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener tareas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
	})
}

// CreateTask crea una nueva tarea
func (tc *TaskController) CreateTask(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	var req CreateTaskRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos de entrada inválidos",
			"details": err.Error(),
		})
		return
	}

	// Establecer status por defecto si no se proporciona
	if req.Status == "" {
		req.Status = models.TaskStatusPending
	}

	task := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		UserID:      userID,
	}

	if err := tc.taskService.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al crear tarea",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tarea creada exitosamente",
		"task":    task,
	})
}

// UpdateTask actualiza una tarea existente
func (tc *TaskController) UpdateTask(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	taskIDStr := c.Param("id")

	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de tarea inválido",
		})
		return
	}

	var req UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Datos de entrada inválidos",
			"details": err.Error(),
		})
		return
	}

	updates := &models.Task{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	task, err := tc.taskService.UpdateTask(uint(taskID), userID, updates)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tarea no encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tarea actualizada exitosamente",
		"task":    task,
	})
}

// DeleteTask elimina una tarea
func (tc *TaskController) DeleteTask(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	taskIDStr := c.Param("id")

	taskID, err := strconv.ParseUint(taskIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID de tarea inválido",
		})
		return
	}

	if err := tc.taskService.DeleteTask(uint(taskID), userID); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Tarea no encontrada",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tarea eliminada exitosamente",
	})
}

// GetTasksSummary obtiene el resumen de tareas usando concurrencia
func (tc *TaskController) GetTasksSummary(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)

	summary, err := tc.taskService.GetTasksSummary(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error al obtener resumen de tareas",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"summary": summary,
	})
}
