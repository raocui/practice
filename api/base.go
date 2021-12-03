package api

import (
	"_/common/conn"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = conn.GetDB()
}
