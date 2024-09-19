package repository

import (
	"assigment3a/service/module/windwater/entity"
	"context"
	"log"
)

func (r WindWaterRepository) UpdateWindWater(ctx context.Context, req entity.WindWater) (lastInsertID uint, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Raw(
		`UPDATE public.wind_water SET wind = ?, water = ?, status_wind = ?, status_water = ?`,
		req.Wind, req.Water, req.StatusWind, req.StatusWater,
	).Scan(&lastInsertID)

	return
}

func (r WindWaterRepository) CreateWindWater(ctx context.Context, req entity.WindWater) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Create(&req)
	if result.Error != nil {
		log.Fatal("Error inserting data:", result.Error)
		return
	}

	return
}

func (r WindWaterRepository) CreateWindWaterLog(ctx context.Context, req entity.WindWaterLog) (err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	result := db.Create(&req)
	if result.Error != nil {
		log.Fatal("Error inserting data:", result.Error)
		return
	}

	return
}

func (r WindWaterRepository) GetIniDataWindWater(ctx context.Context) (data entity.WindWater, err error) {
	db := r.gormDB.GetDB().WithContext(ctx)
	_ = db.Raw(
		`SELECT * FROM public.wind_water limit 1`,
	).Scan(&data)

	return
}
