package services

import (
	"sync"
	"task-manager/internal/config"
	"task-manager/internal/models"
)

type TaskService struct{}

// GetUserTasks obtiene todas las tareas de un usuario
func (s *TaskService) GetUserTasks(userID uint) ([]models.Task, error) {
	db := config.GetDB()
	var tasks []models.Task

	err := db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

// CreateTask crea una nueva tarea
func (s *TaskService) CreateTask(task *models.Task) error {
	db := config.GetDB()
	return db.Create(task).Error
}

// UpdateTask actualiza una tarea existente
func (s *TaskService) UpdateTask(taskID, userID uint, updates *models.Task) (*models.Task, error) {
	db := config.GetDB()
	var task models.Task

	// Verificar que la tarea pertenezca al usuario
	if err := db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		return nil, err
	}

	// Actualizar campos
	if updates.Title != "" {
		task.Title = updates.Title
	}
	if updates.Description != "" {
		task.Description = updates.Description
	}
	if updates.Status != "" {
		task.Status = updates.Status
	}

	if err := db.Save(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

// DeleteTask elimina una tarea
func (s *TaskService) DeleteTask(taskID, userID uint) error {
	db := config.GetDB()
	return db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{}).Error
}

// GetTasksSummary obtiene el resumen de tareas usando goroutines
func (s *TaskService) GetTasksSummary(userID uint) (*models.TaskSummary, error) {
	db := config.GetDB()

	// Crear channels para recibir resultados
	totalCh := make(chan int64, 1)
	completedCh := make(chan int64, 1)
	pendingCh := make(chan int64, 1)
	errorCh := make(chan error, 3)

	var wg sync.WaitGroup

	// Goroutine para contar total de tareas
	wg.Add(1)
	go func() {
		defer wg.Done()
		var count int64
		if err := db.Model(&models.Task{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
			errorCh <- err
			return
		}
		totalCh <- count
	}()

	// Goroutine para contar tareas completadas
	wg.Add(1)
	go func() {
		defer wg.Done()
		var count int64
		if err := db.Model(&models.Task{}).Where("user_id = ? AND status = ?", userID, models.TaskStatusCompleted).Count(&count).Error; err != nil {
			errorCh <- err
			return
		}
		completedCh <- count
	}()

	// Goroutine para contar tareas pendientes
	wg.Add(1)
	go func() {
		defer wg.Done()
		var count int64
		if err := db.Model(&models.Task{}).Where("user_id = ? AND status = ?", userID, models.TaskStatusPending).Count(&count).Error; err != nil {
			errorCh <- err
			return
		}
		pendingCh <- count
	}()

	// Esperar que todas las goroutines terminen
	go func() {
		wg.Wait()
		close(totalCh)
		close(completedCh)
		close(pendingCh)
		close(errorCh)
	}()

	// Verificar errores
	select {
	case err := <-errorCh:
		if err != nil {
			return nil, err
		}
	default:
	}

	// Recoger resultados
	summary := &models.TaskSummary{
		TotalTasks:     <-totalCh,
		CompletedTasks: <-completedCh,
		PendingTasks:   <-pendingCh,
	}

	return summary, nil
}
