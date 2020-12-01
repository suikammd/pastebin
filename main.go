package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"os/signal"
	"syscall"
)

var mysqlConf MysqlConfig
var db *gorm.DB

type Server struct {
	r  *gin.Engine
	db *gorm.DB
}

func init() {
	// TODO: load config from json file
	mysqlConf = MysqlConfig{
		Host:     "127.0.0.1:3306",
		Username: "root",
		Password: "",
		Database: "shorten_url",
	}
	mysqlConf.validate()

	var err error
	db, err = New(mysqlConf)
	if err != nil {
		return
	}
}

func (c *MysqlConfig) validate() {
	if c.Host == "" || c.Username == "" || c.Password == "" || c.Database == "" {
		panic("please check mysql config")
	}
}

func main() {
	r := gin.Default()
	s := Server{
		r:  r,
		db: db,
	}

	// register api
	r.POST("/text", s.PostText)
	r.GET("/link/:short_link", s.GetText)

	go func() {
		s.r.Run()
	}()

	sigquit := make(chan os.Signal, 1)
	signal.Notify(sigquit, os.Interrupt, syscall.SIGTERM)

	_ = <-sigquit
	// TODO close db
}
