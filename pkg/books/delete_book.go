package books

import (
	"docker-go-api/pkg/common/models"
	"docker-go-api/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *handler) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		res := util.BuildErrorResponse("Id not found", util.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, res)
		return
	}

	h.DB.Delete(&book)
	res := util.BuildResponse(true, "Success", nil)
	ctx.JSON(http.StatusOK, res)
}
