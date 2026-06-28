package handlers

import (
	"library/pkg/logs"
	"library/pkg/model/dto"
	"library/pkg/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LibraryHandlers struct {
	ls *service.LibraryService
}

func New(ls *service.LibraryService) *LibraryHandlers {
	return &LibraryHandlers{
		ls: ls,
	}
}

func (lh *LibraryHandlers) Create(ctx *gin.Context) {
	var req dto.CreateBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		logs.LogError(err.Error())
		return
	}

	res, err := lh.ls.Create(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "book with this name already exist",
		})
		logs.LogError(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (lh *LibraryHandlers) AddPage(ctx *gin.Context) {
	name := ctx.Param("book_name")

	res, err := lh.ls.GetByTitle(ctx, &dto.GetBookRequest{Title: name})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unknown book with this name",
		})
		logs.LogError(err.Error())
		return
	}

	req := dto.AddPageRequest{
		BookID: res.Book.ID,
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		logs.LogError(err.Error())
		return
	}

	if err := lh.ls.AddPage(ctx, &req); err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"message": "page with this number already exists",
		})
		logs.LogError(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "page added",
	})
}

func (lh *LibraryHandlers) GetByTitle(ctx *gin.Context) {
	var req dto.GetBookRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		logs.LogError(err.Error())
		return
	}

	res, err := lh.ls.GetByTitle(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "book not found",
		})
		logs.LogError(err.Error())
		return
	}

	ctx.JSON(http.StatusFound, res)
}

func (lh *LibraryHandlers) GetAll(ctx *gin.Context) {
	var req dto.GetAllBookResponse
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		logs.LogError(err.Error())
		return
	}

	res, err := lh.ls.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "not found any books",
		})
		logs.LogError(err.Error())
		return
	}

	ctx.JSON(http.StatusFound, res)
}

func (lh *LibraryHandlers) GetPage(ctx *gin.Context) {
	name := ctx.Param("book_name")
	nStr := ctx.DefaultQuery("p", "1")

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid page number",
		})
		logs.LogError(err.Error())
		return
	}

	res, err := lh.ls.GetByTitle(ctx, &dto.GetBookRequest{Title: name})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unknown book with this name",
		})
		logs.LogError(err.Error())
		return
	}

	resp, err := lh.ls.GetPage(ctx, &dto.GetPageRequest{
		BookID: res.Book.ID,
		Number: uint(n),
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "page not found",
		})
		logs.LogError(err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (lh *LibraryHandlers) Delete(ctx *gin.Context) {
	name := ctx.Param("book_name")

	res, err := lh.ls.GetByTitle(ctx, &dto.GetBookRequest{Title: name})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "unknown book with this name",
		})
		logs.LogError(err.Error())
		return
	}

	if err := lh.ls.Delete(ctx, res.Book.ID); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "book with this name not exists",
		})
		logs.LogError(err.Error())
		return
	}

	ctx.JSON(http.StatusNoContent, "book deleted")
}
