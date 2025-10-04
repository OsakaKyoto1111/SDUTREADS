package controller

import (
	"myapp/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LikeHandler struct {
	service service.LikeService
}

func NewLikeHandler(s service.LikeService) *LikeHandler {
	return &LikeHandler{service: s}
}

func (h *LikeHandler) Toggle(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid post id"})
	}

	userID := c.Get("user_id").(int64)
	action, err := h.service.ToggleLike(userID, postID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": action})
}

func (h *LikeHandler) Count(c echo.Context) error {
	postID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid post id"})
	}

	count, err := h.service.GetLikeCount(postID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"likes": count})
}
