package entity

import "time"

type WindWater struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Wind        int       `json:"wind" gorm:"not null; type: float(8)"`
	Water       int       `json:"water" gorm:"not null; type: float(8)"`
	StatusWind  string    `json:"status_wind" gorm:"not null; type: varchar (100)"`
	StatusWater string    `json:"status_water" gorm:"not null; type: varchar (100)"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null; type: timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null; type: timestamp"`
}

func (WindWater) TableName() string {
	return "public.wind_water"
}

type WindWaterLog struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Wind        int       `json:"wind" gorm:"not null; type: float(8)"`
	Water       int       `json:"water" gorm:"not null; type: float(8)"`
	StatusWind  string    `json:"status_wind" gorm:"not null; type: varchar (100)"`
	StatusWater string    `json:"status_water" gorm:"not null; type: varchar (100)"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null; type: timestamp"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null; type: timestamp"`
}

func (WindWaterLog) TableName() string {
	return "public.wind_water_log"
}
