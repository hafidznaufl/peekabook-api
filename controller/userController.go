package controller

import (
	"net/http"
	"rent-app/context"
	"rent-app/model/web"
	"rent-app/utils/helper"
	"rent-app/utils/res"

	"strings"

	"github.com/labstack/echo/v4"
)

type UserController interface {
	CreateUserController() echo.HandlerFunc
	GetAllUserController() echo.HandlerFunc
}

type UserControllerImpl struct {
	UserContext context.UserContext
}

func NewUserController(userService context.UserContext) UserController {
	return &UserControllerImpl{UserContext: userService}
}

func (UserController *UserControllerImpl) GetAllUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := UserController.UserContext.GetAllUser()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Get Data"))
		}

		results := res.ConvertResponse(data)

		return c.JSON(http.StatusOK, helper.SuccessResponse("Success Retrieve Data", results))
	}
}

func (UserController *UserControllerImpl) CreateUserController() echo.HandlerFunc {
	return func(c echo.Context) error {
		userCreateRequest := web.UserCreateRequest{}
		err := c.Bind(&userCreateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Create Input Error"))
		}

		data, err := UserController.UserContext.CreateUser(userCreateRequest)
		if err != nil {
			if strings.Contains("Email Already Exist", err.Error()) {
				return c.JSON(http.StatusBadRequest, helper.ErrorResponse("Email Already Exist"))
			}

			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Data Error"))
		}

		results := res.UserDomainToUserCreateResponse(data)

		token, err := helper.GenerateToken(&userCreateRequest, uint(data.ID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error Generate Token"))

		}

		results.Token = token

		return c.JSON(http.StatusCreated, helper.SuccessResponse("Success Create User", results))
	}
}
