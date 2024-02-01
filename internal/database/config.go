package database

import "gorm.io/gorm"

type DBManager struct {
	Master *gorm.DB
	Slaves []*gorm.DB
}
