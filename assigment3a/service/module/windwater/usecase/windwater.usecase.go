package usecase

import (
	"assigment3a/service/module/windwater/dto"
	"assigment3a/service/module/windwater/entity"
	"context"

	"github.com/jinzhu/copier"
)

func getWaterStatus(water int) string {
	if water <= 5 {
		return "Aman"
	} else if water >= 6 && water <= 8 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func getWindStatus(wind int) string {
	if wind < 6 {
		return "Aman"
	} else if wind >= 7 && wind <= 15 {
		return "Siaga"
	} else {
		return "Bahaya"
	}
}

func (u *WindWaterUsecase) UpdateData(ctx context.Context, req dto.RequestUpdateWindWater) (respon entity.WindWater, err error) {
	var statusWind = getWindStatus(req.Wind)
	var statusWater = getWaterStatus(req.Water)

	newData := entity.WindWater{
		Wind:        req.Wind,
		Water:       req.Water,
		StatusWind:  statusWind,
		StatusWater: statusWater,
	}
	var newDataLog entity.WindWaterLog
	copier.Copy(&newDataLog, &newData)
	check, err := u.WindWaterRepository.GetIniDataWindWater(ctx)
	if err != nil {
		return
	}
	if check.Id == 0 {
		errCreate := u.WindWaterRepository.CreateWindWater(ctx, newData)
		if errCreate != nil {
			return
		}
		errCreateLog := u.WindWaterRepository.CreateWindWaterLog(ctx, newDataLog)
		if errCreateLog != nil {
			return
		}
	} else {
		_, errUpdate := u.WindWaterRepository.UpdateWindWater(ctx, newData)
		if errUpdate != nil {
			return
		}
		errCreateLog := u.WindWaterRepository.CreateWindWaterLog(ctx, newDataLog)
		if errCreateLog != nil {
			return
		}
	}

	data, err := u.WindWaterRepository.GetIniDataWindWater(ctx)
	respon = data
	return
}
