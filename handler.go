package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io/ioutil"
	"time"
)

func (s Server) PostText(c *gin.Context) {
	pasteText := &PasteText{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("meet unexpected error reading body %s, err is %s", c.Request.Body, err.Error()),
		})
		return
	}

	err = json.Unmarshal(body, pasteText)
	if err != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("meet unexpected error parsing body %s, err is %s", body, err.Error()),
		})
		return
	}

	shortLink := Encode()
	path := fmt.Sprintf("/tmp/%s", shortLink)
	pasteText.Path = path
	paste := &Paste{
		ShortLink:           shortLink,
		CreatedAt:           time.Now(),
		ExpirationInMinutes: 10,
		Path:                path,
		Count:               0,
	}

	res := s.db.Create(&paste)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("meet unexpected error insert body into db, err is %s", err.Error()),
		})
		return
	}

	res = s.db.Create(&pasteText)
	if res.Error != nil {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("meet unexpected error insert body into db, err is %s", err.Error()),
		})
		return
	}

	c.JSON(200, gin.H{
		"short_link": shortLink,
	})
	return
}

func (s Server) GetText(c *gin.Context) {
	shortLink := c.Param("short_link")
	if shortLink == "" {
		c.JSON(400, gin.H{
			"message": "no link specified",
		})
	}

	var paste Paste
	var pasteText PasteText
	res := db.Where("short_link = ?", shortLink).First(&paste)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("no such record with short link %s", shortLink),
		})
	}

	res = db.Where("path = ?", paste.Path).First(&pasteText)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		c.JSON(400, gin.H{
			"message": fmt.Sprintf("no such record with short link %s", shortLink),
		})
	}

	c.JSON(200, gin.H{
		"content": pasteText.Content,
	})
}
