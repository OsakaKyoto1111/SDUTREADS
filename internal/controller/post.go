package controller

import (
	"myapp/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{service: s}
}
func (h *PostHandler) Create(c echo.Context) error {
	var req struct {
		Content string `json:"content"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid json"})
	}

	userID := c.Get("user_id").(int64) // 👈 берём из токена
	if err := h.service.Create(userID, req.Content); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"message": "post created"})
}

func (h *PostHandler) GetAll(c echo.Context) error {
	posts, err := h.service.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, posts)
}

func (h *PostHandler) GetByID(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}
	post, err := h.service.GetByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "post not found"})
	}
	return c.JSON(http.StatusOK, post)
}
func (h *PostHandler) Delete(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid id"})
	}

	userID := c.Get("user_id").(int64)
	if err := h.service.Delete(id, userID); err != nil {
		if err.Error() == "forbidden: you can delete only your own posts" {
			return c.JSON(http.StatusForbidden, echo.Map{"error": err.Error()})
		}
		return c.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "post deleted"})
}
