package handler

import (
	"net/http"
	"painteer/model"
	"painteer/service"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	SignUp(c echo.Context) error
	SignIn(c echo.Context) error
	GetUserByID(c echo.Context) error
}

type AuthHandlerImpl struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandlerImpl {
	return &AuthHandlerImpl{authService: authService}
}

func (h *AuthHandlerImpl) SignUp(c echo.Context) error {
	var req model.CreateUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	user, err := h.authService.RegisterUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, user)
}

func (h *AuthHandlerImpl) Login(c echo.Context) error {
	var req struct {
		AuthId string `json:"auth_id"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	userId, err := h.authService.AuthenticateUser(model.AuthId(req.AuthId))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
	}

	return c.JSON(http.StatusOK, userId)
}

func (h *AuthHandlerImpl) GetUserByID(c echo.Context) error {
	userIdStr := c.Param("user_id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user_id"})
	}

	user, err := h.authService.GetUserByID(model.UserId(userId))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	return c.JSON(http.StatusOK, user)
}
