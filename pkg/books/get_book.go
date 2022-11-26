package books

import (
	"docker-go-api/pkg/common/models"
	"docker-go-api/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		res := util.BuildErrorResponse("id not found", util.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	res := util.BuildResponse(true, "Success", &book)
	ctx.JSON(http.StatusOK, res)
}
