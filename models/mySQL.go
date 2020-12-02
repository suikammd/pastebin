package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func (c MysqlConfig) Validate() {
	if c.Host == "" || c.Username == "" || c.Database == "" {
		panic("please check mysql config")
	}
}

func New(config MysqlConfig) (*gorm.DB, error) {
	// Connect postgres
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local&parseTime=true",
		config.Username, config.Password, config.Host, config.Database)

	db, err := gorm.Open(mysql.Open(connect), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if db.Migrator().HasTable(&Paste{}) {
		db.Migrator().AutoMigrate()
	} else {
		db.Migrator().CreateTable(&Paste{})
	}

	if db.Migrator().HasTable(&PasteText{}) {
		db.Migrator().AutoMigrate()
	} else {
		db.Migrator().CreateTable(&PasteText{})
	}

	// TODO: deal with mysql
	return db, nil
}
