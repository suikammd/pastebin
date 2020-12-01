package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

func (c *MysqlConfig) Validate() {
	// TODO: validate mysql conf
}

var mysqlConf MysqlConfig
var db *gorm.DB

func init()  {
	// TODO: load config from json file
	mysqlConf = MysqlConfig{
		Host: "127.0.0.1:3306",
		Username: "root",
		Password: "",
		Database: "shorten_url",
	}

	var err error
	db, err = New(mysqlConf)
	if err != nil {
		return
	}
}

func main() {
	r := gin.Default()
	r.POST("/text", PostText)
	r.GET("/link/:short_link", GetText)

	go func() {
		r.Run()
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)

	_ = <- sigquit
	// TODO close db
}
