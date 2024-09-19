package dto

type RequestUpdateWindWater struct {
	Wind  int `json:"wind" binding:"required" valid:"required~Data wind wajib diisi"`
	Water int `json:"water" binding:"required" valid:"required~Data water wajib diisi"`
}
