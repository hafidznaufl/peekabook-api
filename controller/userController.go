package controller

import (
	"net/http"
	"peekabook/context"
	"peekabook/model/web"
	"peekabook/utils/helper"
	"peekabook/utils/res"
	"strconv"

	"strings"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	RegisterUserController(ctx echo.Context) error
	LoginUserController(ctx echo.Context) error
	UpdateUserController(ctx echo.Context) error
	GetUserController(ctx echo.Context) error
	GetUsersController(ctx echo.Context) error
	DeleteUserController(ctx echo.Context) error
}

type UserControllerImpl struct {
	UserContext context.UserContext
}

func NewUserController(userContext context.UserContext) UserController {
	return &UserControllerImpl{UserContext: userContext}
}

func (c *UserControllerImpl) RegisterUserController(ctx echo.Context) error {
	userCreateRequest := web.UserCreateRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.UserContext.CreateUser(ctx, userCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Email Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (c *UserControllerImpl) LoginUserController(ctx echo.Context) error {
	userLoginRequest := web.UserLoginRequest{}
	err := ctx.Bind(&userLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := c.UserContext.LoginUser(ctx, userLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Invalid Email or Password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	userLoginResponse := res.UserDomainToUserLoginResponse(response)

	token, err := helper.GenerateToken(&userLoginResponse, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	userLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", userLoginResponse))
}

func (c *UserControllerImpl) GetUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.UserContext.FindById(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "User Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get User Data Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get User Data", response))
}

func (c *UserControllerImpl) GetUsersController(ctx echo.Context) error {
	result, err := c.UserContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Users Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Users Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Users Data Error"))
	}

	response := res.ConvertUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get User Data", response))
}

func (c *UserControllerImpl) DeleteUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.UserContext.DeleteUser(ctx, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "User Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete User Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get User Data", nil))
}

func (c *UserControllerImpl) UpdateUserController(ctx echo.Context) error {
	userId := ctx.Param("id")
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	userUpdateRequest := web.UserUpdateRequest{}
	err = ctx.Bind(&userUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.UserContext.UpdateUser(ctx, userUpdateRequest, userIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "User Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("User Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update User Error"))
	}

	response := res.UserDomaintoUserResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated User", response))
}
