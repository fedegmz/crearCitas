package models

import "gorm.io/gorm"

//citas struct
type Citas struct {
	gorm.Model

	Usuario     string `json:"usuario" gorm:"type:varchar(100);not null"`
	Titulo      string `json:"titulo" gorm:"type:varchar(256)"`
	FechaInicio string `json:"fecha_inicio" gorm:"type:time"`
	FechaFin    string `json:"fecha_fin" gorm:"type:time"`
	AllDay      bool   `json:"all_day" gorm:"type:boolean; default:false"`
}
