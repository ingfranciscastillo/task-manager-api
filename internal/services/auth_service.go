package services

import (
	"errors"
	"task-manager/internal/config"
	"task-manager/internal/models"
	"task-manager/pkg/utils"
)

type AuthService struct{}

// RegisterUser registra un nuevo usuario
func (s *AuthService) RegisterUser(email, password string) (*models.User, error) {
	db := config.GetDB()

	// Verificar si el usuario ya existe
	var existingUser models.User
	if err := db.Where("email = ?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("el usuario ya existe")
	}

	// Crear nuevo usuario
	user := &models.User{
		Email:    email,
		Password: password,
	}

	// Guardar usuario (el hash se hace automáticamente en BeforeCreate)
	if err := db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser autentica un usuario y retorna JWT
func (s *AuthService) LoginUser(email, password string) (string, *models.User, error) {
	db := config.GetDB()

	// Buscar usuario por email
	var user models.User
	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	// Verificar contraseña
	if !user.CheckPassword(password) {
		return "", nil, errors.New("credenciales inválidas")
	}

	// Generar JWT
	token, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}
