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

type BorrowController interface {
	CreateBorrowController(ctx echo.Context) error
	UpdateBorrowController(ctx echo.Context) error
	GetBorrowController(ctx echo.Context) error
	GetBorrowsController(ctx echo.Context) error
	GetBorrowByNameController(ctx echo.Context) error
	DeleteBorrowController(ctx echo.Context) error
}

type BorrowControllerImpl struct {
	BorrowContext context.BorrowContext
}

func NewBorrowController(borrowContext context.BorrowContext) BorrowController {
	return &BorrowControllerImpl{BorrowContext: borrowContext}
}

func (c *BorrowControllerImpl) CreateBorrowController(ctx echo.Context) error {
	borrowCreateRequest := web.BorrowCreateRequest{}
	err := ctx.Bind(&borrowCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.BorrowContext.CreateBorrow(ctx, borrowCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Borrow Error"))
	}

	response := res.BorrowDomaintoBorrowResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Borrow Data", response))
}

func (c *BorrowControllerImpl) GetBorrowController(ctx echo.Context) error {
	borrowId := ctx.Param("id")
	borrowIdInt, err := strconv.Atoi(borrowId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.BorrowContext.FindById(ctx, borrowIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Borrow Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Borrow Data Error"))
	}

	response := res.BorrowDomaintoBorrowResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Borrow Data", response))
}

func (c *BorrowControllerImpl) GetBorrowsController(ctx echo.Context) error {
	result, err := c.BorrowContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Borrows Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrows Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Borrows Data Error"))
	}

	response := res.ConvertBorrowResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Borrows Data", response))
}

func (c *BorrowControllerImpl) GetBorrowByNameController(ctx echo.Context) error {
	borrowName := ctx.Param("name")

	result, err := c.BorrowContext.FindByName(ctx, borrowName)
	if err != nil {
		if strings.Contains(err.Error(), "Borrow Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Borrow Data By Name Error"))
	}

	response := res.BorrowDomaintoBorrowResponse(result)
	fmt.Println(response)
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Borrow Data By Name", response))
}

func (c *BorrowControllerImpl) UpdateBorrowController(ctx echo.Context) error {
	borrowId := ctx.Param("id")
	borrowIdInt, err := strconv.Atoi(borrowId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	borrowUpdateRequest := web.BorrowUpdateRequest{}
	err = ctx.Bind(&borrowUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Client Input"))
	}

	result, err := c.BorrowContext.UpdateBorrow(ctx, borrowUpdateRequest, borrowIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "Borrow Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Borrow Error"))
	}

	response := res.BorrowDomaintoBorrowResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Updated Borrow Data", response))
}

func (c *BorrowControllerImpl) DeleteBorrowController(ctx echo.Context) error {
	borrowId := ctx.Param("id")
	borrowIdInt, err := strconv.Atoi(borrowId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.BorrowContext.DeleteBorrow(ctx, borrowIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Borrow Not Found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Borrow Data Error"))
	}

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Deleted Borrow Data", nil))
}
