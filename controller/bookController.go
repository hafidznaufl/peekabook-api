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

type BookController interface {
	CreateBookController(ctx echo.Context) error
	UpdateBookController(ctx echo.Context) error
	GetBookController(ctx echo.Context) error
	GetBooksController(ctx echo.Context) error
	GetBookByNameController(ctx echo.Context) error
	DeleteBookController(ctx echo.Context) error
}

type BookControllerImpl struct {
	BookContext context.BookContext
}

func NewBookController(bookContext context.BookContext) BookController {
	return &BookControllerImpl{BookContext: bookContext}
}

func (c *BookControllerImpl) CreateBookController(ctx echo.Context) error {
	bookCreateRequest := web.BookCreateRequest{}
	err := ctx.Bind(&bookCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.BookContext.CreateBook(ctx, bookCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Book Error"))
	}

	response := res.BookDomaintoBookResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Book", response))
}

func (c *BookControllerImpl) GetBookController(ctx echo.Context) error {
	bookId := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.BookContext.FindById(ctx, bookIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Book Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Book Data Error"))
	}

	response := res.BookDomaintoBookResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Book Data", response))
}

func (c *BookControllerImpl) GetBooksController(ctx echo.Context) error {
	result, err := c.BookContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Books Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Books Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Books Data Error"))
	}

	response := res.ConvertBookResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Book Data", response))
}

func (c *BookControllerImpl) GetBookByNameController(ctx echo.Context) error {
	bookName := ctx.Param("name")

	result, err := c.BookContext.FindByName(ctx, bookName)
	if err != nil {
		if strings.Contains(err.Error(), "Book Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Book Data By Name Error"))
	}

	response := res.BookDomaintoBookResponse(result)
	fmt.Println(response)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Book Data By Name", response))
}

func (c *BookControllerImpl) UpdateBookController(ctx echo.Context) error {
	bookId := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	bookUpdateRequest := web.BookUpdateRequest{}
	err = ctx.Bind(&bookUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.BookContext.UpdateBook(ctx, bookUpdateRequest, bookIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Book Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Book Error"))
	}

	response := res.BookDomaintoBookResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Book", response))
}

func (c *BookControllerImpl) DeleteBookController(ctx echo.Context) error {
	bookId := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.BookContext.DeleteBook(ctx, bookIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Book Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Book Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Get Book Data", nil))
}
