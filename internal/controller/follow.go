package controller

import (
	"myapp/internal/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FollowHandler struct {
	service service.FollowService
}

func NewFollowHandler(s service.FollowService) *FollowHandler {
	return &FollowHandler{service: s}
}

func (h *FollowHandler) ToggleFollow(c echo.Context) error {
	followingID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	followerID := c.Get("user_id").(int64)
	status, err := h.service.ToggleFollow(followerID, followingID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"status": status})
}

func (h *FollowHandler) Feed(c echo.Context) error {
	userID := c.Get("user_id").(int64)
	posts, err := h.service.GetFeedPosts(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"feed": posts})
}
