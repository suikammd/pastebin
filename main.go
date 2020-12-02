package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ini/ini"
	"github.com/suikammd/shorten-url/models"
	"gorm.io/gorm"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	mysqlConf models.MysqlConfig
	db        *gorm.DB
	cfg       *ini.File
)

type Server struct {
	r  *gin.Engine
	db *gorm.DB
}

func init() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	loadDB()
}

func loadDB() {
	dbConf, err := cfg.GetSection("database")
	if err != nil {
		log.Fatalf("Fail to get section 'database': %v", err)
	}

	host := dbConf.Key("HOST").MustString("127.0.0.1")
	username := dbConf.Key("USERNAME").MustString("root")
	password := dbConf.Key("PASSWORD").MustString("")
	database := dbConf.Key("DATABASE").MustString("shorten_url")
	mysqlConf = models.MysqlConfig{
		Host:     host,
		Username: username,
		Password: password,
		Database: database,
	}

	mysqlConf.validate()

	db, err = models.New(mysqlConf)
	if err != nil {
		return
	}
}

func (c *models.MysqlConfig) validate() {
	if c.Host == "" || c.Username == "" || c.Database == "" {
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
