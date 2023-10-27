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

type BookController interface {
	CreateBookController(ctx echo.Context) error
	UpdateBookController(ctx echo.Context) error
	GetBookController(ctx echo.Context) error
	GetBooksController(ctx echo.Context) error
	GetBookByTitleController(ctx echo.Context) error
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
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Book Error"))
	}

	response := res.CreateBookDomaintoBookResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Book Data", response))
}

func (c *BookControllerImpl) GetBookController(ctx echo.Context) error {
	bookId := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.BookContext.FindById(ctx, bookIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "book not found") {
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
		if strings.Contains(err.Error(), "books not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Books Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Books Data Error"))
	}

	response := res.ConvertBookResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Book Data", response))
}

func (c *BookControllerImpl) GetBookByTitleController(ctx echo.Context) error {
	bookTitle := ctx.Param("name")

	result, err := c.BookContext.FindByTitle(ctx, bookTitle)
	if err != nil {
		if strings.Contains(err.Error(), "book not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Book Data By Title Error"))
	}

	response := res.BookDomaintoBookResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Book Data By Title", response))
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
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "book not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Book Error"))
	}

	response := res.UpdateBookDomaintoBookResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Updated Book Data", response))
}

func (c *BookControllerImpl) DeleteBookController(ctx echo.Context) error {
	bookId := ctx.Param("id")
	bookIdInt, err := strconv.Atoi(bookId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.BookContext.DeleteBook(ctx, bookIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "book not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Book Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Book Data Error"))
	}

	return ctx.JSON(http.StatusNoContent, helper.SuccessResponse("Successfully Deleted Book Data", nil))
}
