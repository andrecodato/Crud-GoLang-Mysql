package models

import (
	"github.com/andrecodato/go-hardwarestore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Hardware struct {
	gorm.Model
	Name   string `json:"name"`
	Brand  string `json:"brand"`
	Serial string `json:"serial"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Hardware{})
}

func (h *Hardware) CreateHardware() *Hardware {
	db.NewRecord(h)
	db.Create(&h)
	return h
}

func GetAllHardwares() []Hardware {
	var Hardwares []Hardware
	db.Find(&Hardwares)
	return Hardwares
}

func GetHardwareById(Id int64) (*Hardware, *gorm.DB) {
	var getHardware Hardware
	db := db.Where("ID=?", Id).Find(&getHardware)
	return &getHardware, db
}

func DeleteHardware(Id int64) Hardware {
	var hardware Hardware
	db.Where("ID=?", Id).Delete(hardware)
	return hardware
}
