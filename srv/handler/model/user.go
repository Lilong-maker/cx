package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	Name  string  `grom:"type:varchar(30)"`
	Price float64 `grom:"type:decimal(10,2)"`
	Num   int     `grom:"type:int"`
}

func (o *Order) FindOrder(db *gorm.DB, name string) error {
	return db.Debug().Where("name = ?", name).Find(&o).Error
}

func (o *Order) OrderAdd(db *gorm.DB) error {
	return db.Debug().Create(&o).Error
}
