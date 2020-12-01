package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func PostText(c *gin.Context) {
	pasteText := &PasteText{}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println("err is " + err.Error())
		return
	}

	err = json.Unmarshal(body, pasteText)
	if err != nil {
		fmt.Println("err is " + err.Error())
		return
	}
	c.JSON(200, gin.H{
		"message": pasteText.Content,
	})
}

func GetText(c *gin.Context) {
	shortLink := c.Param("short_link")
	if shortLink == "" {
		c.JSON(400, gin.H{
			"message": "no link specified",
		})
	}

	c.JSON(200, gin.H{
		"message": shortLink,
		"query": c.Query("ip"),
	})
}
