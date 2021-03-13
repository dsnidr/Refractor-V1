package api

import (
	"github.com/labstack/echo/v4"
	"github.com/sniddunc/refractor/internal/params"
	"github.com/sniddunc/refractor/pkg/config"
	"github.com/sniddunc/refractor/pkg/jwt"
	"github.com/sniddunc/refractor/refractor"
	"net/http"
	"strconv"
)

type userHandler struct {
	service refractor.UserService
}

func NewUserHandler(userService refractor.UserService) refractor.UserHandler {
	return &userHandler{
		service: userService,
	}
}

func (h *userHandler) GetOwnUserInfo(c echo.Context) error {
	claims := c.Get("claims").(*jwt.Claims)

	userInfo, res := h.service.GetUserInfo(claims.UserID)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: userInfo,
	})
}

func (h *userHandler) ChangeUserPassword(c echo.Context) error {
	body := params.ChangeUserPassword{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	_, res := h.service.ChangeUserPassword(claims.UserID, body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
	})
}

func (h *userHandler) GetAllUsers(c echo.Context) error {
	users, res := h.service.GetAllUsers()
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Payload: users,
	})
}

func (h *userHandler) CreateUser(c echo.Context) error {
	body := params.CreateUserParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	// Create the user
	_, res := h.service.CreateUser(body)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
		Errors:  res.ValidationErrors,
	})
}

func (h *userHandler) ActivateUser(c echo.Context) error {
	idString := c.Param("id")

	userID, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageInvalidIDProvided,
		})
	}

	// Activate user
	_, res := h.service.UpdateUser(userID, refractor.UpdateArgs{
		"Activated": true,
	})

	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
	})
}

func (h *userHandler) DeactivateUser(c echo.Context) error {
	idString := c.Param("id")

	userID, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageInvalidIDProvided,
		})
	}

	// Deactivate user
	_, res := h.service.UpdateUser(userID, refractor.UpdateArgs{
		"Activated": false,
	})

	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
	})
}

func (h *userHandler) ForcePasswordChange(c echo.Context) error {
	idString := c.Param("id")

	userID, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Success: false,
			Message: config.MessageInvalidIDProvided,
		})
	}

	claims := c.Get("claims").(*jwt.Claims)

	setterMeta := &params.UserMeta{
		UserID:      claims.UserID,
		Permissions: claims.Permissions,
	}

	res := h.service.ForceUserPasswordChange(userID, setterMeta)
	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
	})
}

func (h *userHandler) SetUserPassword(c echo.Context) error {
	body := params.SetUserPasswordParams{}
	if ok := ValidateRequest(&body, c); !ok {
		return nil
	}

	claims := c.Get("claims").(*jwt.Claims)

	// Update body to include the setting user's details
	body.UserMeta = &params.UserMeta{
		UserID:      claims.UserID,
		Permissions: claims.Permissions,
	}

	_, res := h.service.SetUserPassword(body)

	return c.JSON(res.StatusCode, Response{
		Success: res.Success,
		Message: res.Message,
	})
}
