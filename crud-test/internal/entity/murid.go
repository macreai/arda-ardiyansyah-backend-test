package entity

import "time"

type Murid struct {
	ID         int       `gorm:"primaryKey"`
	Name       string    `gorm:"column:name"`
	TimeCreate time.Time `gorm:"column:time_create;autoCreateTime"`
}

func (m *Murid) TableName() string {
	return "murid"
}
