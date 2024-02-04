package controllers

import (
	"github.com/agung96tm/miblog/api/dto"
	"github.com/agung96tm/miblog/api/models"
	"github.com/agung96tm/miblog/api/services"
	"github.com/agung96tm/miblog/constants"
	"github.com/agung96tm/miblog/pkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct {
	authService services.AuthService
	userService services.UserService
}

func NewUserController(authService services.AuthService, userService services.UserService) UserController {
	return UserController{
		authService: authService,
		userService: userService,
	}
}

// Me godoc
//
//	@Summary		Get data Logged-in user
//	@Description	Get data Logged-in user
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Router			/me [get]
//	@Success		200  {object}  response.Response{data=dto.MeResponse}  "ok"
func (c UserController) Me(ctx echo.Context) error {
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	return ctx.JSON(http.StatusOK, dto.MeResponse{
		Name:  user.Name,
		Email: user.Email,
	})
}

// MePassword godoc
//
//	@Summary		Get data Logged-in user
//	@Description	Get data Logged-in user
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.MePasswordRequest true "Post"
//	@Router			/me/password [post]
//	@Success		200  {object}  response.Response{}  "ok"
func (c UserController) MePassword(ctx echo.Context) error {
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	var passwordReq dto.MePasswordRequest
	if err := ctx.Bind(&passwordReq); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	err := c.userService.MeUpdatePassword(user, &passwordReq)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return response.Response{
		Code: http.StatusAccepted,
	}.JSON(ctx)
}

// MeUpdate godoc
//
//	@Summary		Update data Logged-in user
//	@Description	Update data Logged-in user
//	@Tags			user
//	@Accept			application/json
//	@Produce		application/json
//	@Param 			data body dto.MeUpdateRequest true "Post"
//	@Router			/me [patch]
//	@Success		200  {object}  response.Response{data=dto.MeResponse}  "ok"
func (c UserController) MeUpdate(ctx echo.Context) error {
	user := ctx.Get(constants.CurrentUser).(*models.User)
	if user.IsAnonymous() {
		return response.Response{
			Code: http.StatusUnauthorized,
		}.JSON(ctx)
	}

	var meRequest dto.MeUpdateRequest
	if err := ctx.Bind(&meRequest); err != nil {
		return response.Response{
			Error: err,
		}.JSONValidationError(ctx)
	}

	err := c.userService.MeUpdate(user, &meRequest)
	if err != nil {
		return response.Response{
			Code:    http.StatusBadRequest,
			Message: err,
		}.JSON(ctx)
	}

	return ctx.JSON(http.StatusOK, dto.MeResponse{
		Name:  user.Name,
		Email: user.Email,
	})
}