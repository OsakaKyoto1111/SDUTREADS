package controller

import (
	"myapp/internal/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(r repository.UserRepository) *UserHandler {
	return &UserHandler{repo: r}
}
func (h *UserHandler) Me(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	user, err := h.repo.FindByID(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}
