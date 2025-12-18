package handler

import (
	"errors"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"

	"github.com/example/user-age-api/internal/models"
	"github.com/example/user-age-api/internal/service"
)

type UserHandler struct {
	service service.UserService
	logger  *zap.Logger
}

func NewUserHandler(s service.UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: s,
		logger:  logger,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.service.CreateUser(c.Context(), req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return fiber.NewError(fiber.StatusBadRequest, "validation failed")
		}
		h.logger.Error("failed to create user", zap.Error(err))
		return fiber.NewError(fiber.StatusInternalServerError, "failed to create user")
	}

	resp := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.DOB.Format("2006-01-02"),
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	user, err := h.service.GetUser(c.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		h.logger.Error("failed to get user", zap.Error(err))
		return fiber.NewError(fiber.StatusInternalServerError, "failed to get user")
	}
///agee is calculated here 
	age := models.CalculateAge(user.DOB, time.Now().UTC())
	resp := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.DOB.Format("2006-01-02"),
		Age:  age,
	}

	return c.JSON(resp)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		h.logger.Error("failed to list users", zap.Error(err))
		return fiber.NewError(fiber.StatusInternalServerError, "failed to list users")
	}

	now := time.Now().UTC()
	resp := make([]models.UserResponse, 0, len(users))
	for _, u := range users {
		resp = append(resp, models.UserResponse{
			ID:   u.ID,
			Name: u.Name,
			DOB:  u.DOB.Format("2006-01-02"),
			Age:  models.CalculateAge(u.DOB, now),
		})
	}

	return c.JSON(resp)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	user, err := h.service.UpdateUser(c.Context(), id, req)
	if err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			return fiber.NewError(fiber.StatusBadRequest, "validation failed")
		}
		if errors.Is(err, pgx.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		h.logger.Error("failed to update user", zap.Error(err))
		return fiber.NewError(fiber.StatusInternalServerError, "failed to update user")
	}

	resp := models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.DOB.Format("2006-01-02"),
	}

	return c.JSON(resp)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	if err := h.service.DeleteUser(c.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return fiber.NewError(fiber.StatusNotFound, "user not found")
		}
		h.logger.Error("failed to delete user", zap.Error(err))
		return fiber.NewError(fiber.StatusInternalServerError, "failed to delete user")
	}

	return c.SendStatus(fiber.StatusNoContent)
}


