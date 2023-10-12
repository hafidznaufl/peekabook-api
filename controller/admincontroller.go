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

type AdminController interface {
	RegisterAdminController(ctx echo.Context) error
	LoginAdminController(ctx echo.Context) error
	UpdateAdminController(ctx echo.Context) error
	GetAdminController(ctx echo.Context) error
	GetAdminsController(ctx echo.Context) error
	DeleteAdminController(ctx echo.Context) error
}

type AdminControllerImpl struct {
	AdminContext context.AdminContext
}

func NewAdminController(adminContext context.AdminContext) AdminController {
	return &AdminControllerImpl{AdminContext: adminContext}
}

func (c *AdminControllerImpl) RegisterAdminController(ctx echo.Context) error {
	adminCreateRequest := web.AdminCreateRequest{}
	err := ctx.Bind(&adminCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AdminContext.CreateAdmin(ctx, adminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "Email Already Exist") {
			return ctx.JSON(http.StatusConflict, helper.ErrorResponse("Email Already Exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign Up Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign Up", response))
}

func (c *AdminControllerImpl) LoginAdminController(ctx echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}
	err := ctx.Bind(&adminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	response, err := c.AdminContext.LoginAdmin(ctx, adminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Invalid Email or Password") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Email or Password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Sign In Error"))
	}

	adminLoginResponse := res.AdminDomainToAdminLoginResponse(response)

	token, err := helper.GenerateTokenAdmin(&adminLoginResponse, uint(response.ID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Generate JWT Error"))
	}

	adminLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Sign In", adminLoginResponse))
}

func (c *AdminControllerImpl) GetAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.AdminContext.FindById(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Admin Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admin Data Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Admin Data", response))
}

func (c *AdminControllerImpl) GetAdminsController(ctx echo.Context) error {
	result, err := c.AdminContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Admins Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admins Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Admins Data Error"))
	}

	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Admin Data", response))
}

func (c *AdminControllerImpl) DeleteAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.AdminContext.DeleteAdmin(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Admin Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Admin Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Admin Data", nil))
}

func (c *AdminControllerImpl) UpdateAdminController(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	adminUpdateRequest := web.AdminUpdateRequest{}
	err = ctx.Bind(&adminUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AdminContext.UpdateAdmin(ctx, adminUpdateRequest, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Admin Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Admin Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Admin Error"))
	}

	response := res.AdminDomaintoAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Admin", response))
}
