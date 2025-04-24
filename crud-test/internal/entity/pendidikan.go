package entity

import "time"

type Pendidikan struct {
	ID         int       `gorm:"primaryKey"`
	IDMurid    int       `gorm:"column:id_murid"`
	Murid      Murid     `gorm:"foreignKey:IDMurid;references:id"`
	Status     string    `gorm:"column:status"`
	TimeCreate time.Time `gorm:"time_create"`
}
