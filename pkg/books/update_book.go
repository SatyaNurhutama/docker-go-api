package books

import (
	"docker-go-api/pkg/books/dto"
	"docker-go-api/pkg/common/models"
	"docker-go-api/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) UpdateBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	body := dto.BookRequest{}

	if err := ctx.ShouldBind(&body); err != nil {
		res := util.BuildErrorResponse(err.Error(), util.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		res := util.BuildErrorResponse("id not found", util.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	book.Title = body.Title
	book.Description = body.Description
	book.Author = body.Author

	h.DB.Save(&book)
	res := util.BuildResponse(true, "Success", &book)
	ctx.JSON(http.StatusOK, res)
}
