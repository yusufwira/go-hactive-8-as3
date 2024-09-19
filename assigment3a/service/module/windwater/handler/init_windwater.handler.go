package handler

import (
	usecase "assigment3a/service/module/windwater/usecase"
)

type WindWaterHandler struct {
	WindWaterUsecase usecase.IWindWaterUsecase
}

func InitWindWaterHandler(windwaterUsecase usecase.IWindWaterUsecase) *WindWaterHandler {
	return &WindWaterHandler{
		WindWaterUsecase: windwaterUsecase,
	}
}
