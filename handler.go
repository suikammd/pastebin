package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/suikammd/shorten-url/models"
	"github.com/suikammd/shorten-url/pkg/e"
	"github.com/suikammd/shorten-url/pkg/util"
	"gorm.io/gorm"
	"io/ioutil"
	"time"
)

func (s Server) PostText(c *gin.Context) {
	pasteText := &models.PasteText{}
	message := ""
	defer func() {
		if message != "" {
			c.JSON(400, gin.H{
				"message": message,
			})
			return
		}
	}()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		message = e.GetMsg(e.ReadRequestError)
		return
	}

	err = json.Unmarshal(body, pasteText)
	if err != nil {
		message = e.GetMsg(e.ParseRequestError)
		return
	}

	shortLink := util.Encode()
	path := fmt.Sprintf("/tmp/%s", shortLink)
	pasteText.Path = path
	paste := &models.Paste{
		ShortLink:           shortLink,
		CreatedAt:           time.Now(),
		ExpirationInMinutes: 10,
		Path:                path,
		Count:               0,
	}

	res := s.db.Create(&paste)
	if res.Error != nil {
		message = e.GetMsg(e.CreateError)
		return
	}

	res = s.db.Create(&pasteText)
	if res.Error != nil {
		message = e.GetMsg(e.CreateError)
		return
	}

	c.JSON(200, gin.H{
		"short_link": shortLink,
	})
	return
}

func (s Server) GetText(c *gin.Context) {
	shortLink := c.Param("short_link")
	message := ""
	defer func() {
		if message != "" {
			c.JSON(400, gin.H{
				"message": message,
			})
			return
		}
	}()
	if shortLink == "" {
		message = e.GetMsg(e.NoShortLinkError)
		return
	}

	var paste models.Paste
	var pasteText models.PasteText
	res := db.Where("short_link = ?", shortLink).First(&paste)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		message = e.GetMsg(e.NOTFOUND)
		return
	}

	res = db.Where("path = ?", paste.Path).First(&pasteText)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		message = e.GetMsg(e.NOTFOUND)
		return
	}

	// check expiration
	if time.Now().Sub(paste.CreatedAt) > time.Minute * time.Duration(paste.ExpirationInMinutes) {
		// delete paste & paste text
		db.Delete(&paste)
		db.Delete(&pasteText)
		message = e.GetMsg(e.EXPIRED)
		return
	}

	c.JSON(200, gin.H{
		"content": pasteText.Content,
	})
}
