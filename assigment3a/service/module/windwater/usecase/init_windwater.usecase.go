package usecase

import (
	"assigment3a/connection"
	"assigment3a/service/module/windwater/dto"
	"assigment3a/service/module/windwater/entity"
	repo "assigment3a/service/module/windwater/repository"
	"context"
)

type WindWaterUsecase struct {
	WindWaterRepository repo.IWindWaterRepository
	GormDB              *connection.GormDB
}

func InitWindWaterUsecase(windWaterRepository repo.IWindWaterRepository, gormDB *connection.GormDB) IWindWaterUsecase {
	return &WindWaterUsecase{
		WindWaterRepository: windWaterRepository,
		GormDB:              gormDB,
	}
}

type IWindWaterUsecase interface {
	UpdateData(ctx context.Context, req dto.RequestUpdateWindWater) (respon entity.WindWater, err error)
}
