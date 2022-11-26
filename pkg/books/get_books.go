package books

import (
	"docker-go-api/pkg/common/models"
	"docker-go-api/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBooks(ctx *gin.Context) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		res := util.BuildErrorResponse(result.Error.Error(), util.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := util.BuildResponse(true, "Success", &books)
	ctx.JSON(http.StatusOK, res)
}
