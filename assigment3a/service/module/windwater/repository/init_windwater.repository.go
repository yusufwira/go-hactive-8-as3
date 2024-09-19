package repository

import (
	"assigment3a/connection"
	"assigment3a/service/module/windwater/entity"
	"context"
)

type WindWaterRepository struct {
	gormDB *connection.GormDB
}

func InitWindWaterRepository(gormDB *connection.GormDB) IWindWaterRepository {
	return &WindWaterRepository{
		gormDB: gormDB,
	}
}

type IWindWaterRepository interface {
	GetIniDataWindWater(ctx context.Context) (data entity.WindWater, err error)
	CreateWindWater(ctx context.Context, req entity.WindWater) (err error)
	CreateWindWaterLog(ctx context.Context, req entity.WindWaterLog) (err error)
	UpdateWindWater(ctx context.Context, req entity.WindWater) (lastInsertID uint, err error)
}
