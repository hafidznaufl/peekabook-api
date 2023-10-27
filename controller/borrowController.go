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

type BorrowController interface {
	CreateBorrowController(ctx echo.Context) error
	ReturnBorrowController(ctx echo.Context) error
	UpdateBorrowController(ctx echo.Context) error
	GetBorrowController(ctx echo.Context) error
	GetBorrowsByUserNameController(ctx echo.Context) error
	GetBorrowsController(ctx echo.Context) error
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
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))

		}

		if strings.Contains(err.Error(), "unavailable") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("The Book is Unavailable for Borrowing"))

		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Create Borrow Error"))
	}

	response := res.CreateBorrowDomaintoBorrowResponse(result)

	return ctx.JSON(http.StatusCreated, helper.SuccessResponse("Successfully Create Borrow Data", response))
}

func (controller *BorrowControllerImpl) ReturnBorrowController(ctx echo.Context) error {
	// Mendapatkan borrowID dari URL
	borrowId := ctx.Param("id")
	borrowIdInt, err := strconv.Atoi(borrowId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	// Panggil fungsi context untuk mengembalikan buku menggunakan borrowID dari URL
	result, err := controller.BorrowContext.ReturnBorrow(ctx, borrowIdInt)
	if err != nil {
		// Mengatasi kesalahan yang mungkin terjadi selama proses pengembalian
		if err.Error() == "unavailable" {
			// Mengembalikan respons "Unavailable" jika buku tidak dapat dikembalikan
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("The book is unavailable for returning"))
		}
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Error returning the book"))
	}

	response := res.ReturnBorrowDomaintoBorrowResponse(result)

	// Mengembalikan respons sukses jika pengembalian buku berhasil
	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Book returned successfully", response))
}

func (c *BorrowControllerImpl) GetBorrowController(ctx echo.Context) error {
	borrowId := ctx.Param("id")
	borrowIdInt, err := strconv.Atoi(borrowId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	result, err := c.BorrowContext.FindById(ctx, borrowIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "borrow not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Borrow Data Error"))
	}

	response := res.BorrowDomaintoBorrowResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Borrow Data", response))
}

func (c *BorrowControllerImpl) GetBorrowsByUserNameController(ctx echo.Context) error {
	userName := ctx.Param("name")

	result, err := c.BorrowContext.FindBorrowsByUserName(ctx, userName)
	if err != nil {
		if strings.Contains(err.Error(), "borrows not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrows Not Found for User"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get Borrow Data Error"))
	}

	response := res.ConvertBorrowResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get Borrows for User", response))
}

func (c *BorrowControllerImpl) GetBorrowsController(ctx echo.Context) error {
	result, err := c.BorrowContext.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "borrows not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrows Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Get All Borrows Data Error"))
	}

	response := res.ConvertBorrowResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Get All Borrows Data", response))
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
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helper.ErrorResponse("Invalid Validation"))
		}

		if strings.Contains(err.Error(), "borrow not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Update Borrow Error"))
	}

	response := res.BorrowDomaintoBorrowResponse(result)

	return ctx.JSON(http.StatusOK, helper.SuccessResponse("Successfully Updated Borrow Data", response))
}

func (c *BorrowControllerImpl) DeleteBorrowController(ctx echo.Context) error {
	borrowId := ctx.Param("id")
	borrowIdInt, err := strconv.Atoi(borrowId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Invalid Param Id"))
	}

	err = c.BorrowContext.DeleteBorrow(ctx, borrowIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "borrow not found") {
			return ctx.JSON(http.StatusNotFound, helper.ErrorResponse("Borrow Not Found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helper.ErrorResponse("Delete Borrow Data Error"))
	}

	return ctx.JSON(http.StatusNoContent, helper.SuccessResponse("Successfully Deleted Borrow Data", nil))
}
