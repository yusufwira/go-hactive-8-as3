package handler

import (
	"assigment3a/service/module/windwater/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h WindWaterHandler) Update(ctx *gin.Context) {
	req := dto.RequestUpdateWindWater{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Error",
			"data":    err.Error(),
		})
		return
	}

	res, err := h.WindWaterUsecase.UpdateData(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error",
			"data":    err,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    res,
	})
}
