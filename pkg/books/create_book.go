package books

import (
	"docker-go-api/pkg/books/dto"
	"docker-go-api/pkg/common/models"
	"docker-go-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) CreateBook(ctx *gin.Context) {
	body := dto.BookRequest{}

	if err := ctx.ShouldBind(&body); err != nil {
		response := util.BuildErrorResponse(err.Error(), util.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var book models.Book

	book.Title = body.Title
	book.Description = body.Description
	book.Author = body.Author

	if result := h.DB.Create(&book); result.Error != nil {
		response := util.BuildErrorResponse(result.Error.Error(), util.EmptyObj{})
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	response := util.BuildResponse(true, "Success", &book)
	ctx.JSON(http.StatusOK, response)
}
