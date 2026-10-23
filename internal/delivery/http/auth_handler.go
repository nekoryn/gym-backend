package http

import (
	"gym-backend/internal/domain"
	"gym-backend/internal/service"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

// Register godoc
// @Summary Регистрация нового пользователя
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.AuthDTO true "Данные для регистрации"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /api/auth/register [post]
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var dto domain.AuthDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат запроса"})
	}

	if err := h.authService.Register(dto.Email, dto.Password); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Пользователь с таким email уже существует"})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Пользователь успешно зарегистрирован"})
}

// Login godoc
// @Summary Вход в систему и получение JWT токена
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body domain.AuthDTO true "Данные для входа"
// @Success 200 {object} domain.AuthResponse
// @Failure 401 {object} map[string]string
// @Router /api/auth/login [post]
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var dto domain.AuthDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Неверный формат запроса"})
	}

	token, err := h.authService.Login(dto.Email, dto.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(domain.AuthResponse{Token: token})
}