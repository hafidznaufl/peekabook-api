package controller

import (
	"fmt"
	"net/http"
	"peekabook/context"
	"peekabook/model/web"
	"peekabook/utils/helper"
	"peekabook/utils/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthorController interface {
	CreateAuthorController(ctx echo.Context) error
	UpdateAuthorController(ctx echo.Context) error
	GetAuthorController(ctx echo.Context) error
	GetAuthorsController(ctx echo.Context) error
	GetAuthorByNameController(ctx echo.Context) error
	DeleteAuthorController(ctx echo.Context) error
}

type AuthorControllerImpl struct {
	AuthorContext context.AuthorContext
}

func NewAuthorController(authorContext context.AuthorContext) AuthorController {
	return &AuthorControllerImpl{AuthorContext: authorContext}
}

func (c *AuthorControllerImpl) CreateAuthorController(ctx echo.Context) error {
	authorCreateRequest := web.AuthorCreateRequest{}
	err := ctx.Bind(&authorCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AuthorContext.CreateAuthor(ctx, authorCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Author Error"))
	}

	response := res.AuthorDomaintoAuthorResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Author Data", response))
}

func (c *AuthorControllerImpl) GetAuthorController(ctx echo.Context) error {
	authorId := ctx.Param("id")
	authorIdInt, err := strconv.Atoi(authorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.AuthorContext.FindById(ctx, authorIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Author Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Author Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Author Data Error"))
	}

	response := res.AuthorDomaintoAuthorResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Author Data", response))
}

func (c *AuthorControllerImpl) GetAuthorsController(ctx echo.Context) error {
	result, err := c.AuthorContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Authors Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Authors Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Authors Data Error"))
	}

	response := res.ConvertAuthorResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Author Data", response))
}

func (c *AuthorControllerImpl) GetAuthorByNameController(ctx echo.Context) error {
	authorName := ctx.Param("name")

	result, err := c.AuthorContext.FindByName(ctx, authorName)
	if err != nil {
		if strings.Contains(err.Error(), "Author Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Author Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Author Data By Name Error"))
	}

	response := res.AuthorDomaintoAuthorResponse(result)
	fmt.Println(response)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Author Data By Name", response))
}

func (c *AuthorControllerImpl) UpdateAuthorController(ctx echo.Context) error {
	authorId := ctx.Param("id")
	authorIdInt, err := strconv.Atoi(authorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	authorUpdateRequest := web.AuthorUpdateRequest{}
	err = ctx.Bind(&authorUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.AuthorContext.UpdateAuthor(ctx, authorUpdateRequest, authorIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Author Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Author Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Author Error"))
	}

	response := res.AuthorDomaintoAuthorResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Author Data", response))
}

func (c *AuthorControllerImpl) DeleteAuthorController(ctx echo.Context) error {
	authorId := ctx.Param("id")
	authorIdInt, err := strconv.Atoi(authorId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.AuthorContext.DeleteAuthor(ctx, authorIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Author Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Author Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Author Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Author Data", nil))
}
