package filserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	var body CreateBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var filePath = FileDir + body.Name

	err := ioutil.WriteFile(filePath, []byte(body.Content), os.ModePerm)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{
		"stuatus": "ok",
	})
}

func Read(c *gin.Context) {
	filePath := FileDir + c.Param("name")

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件不存在"})
		return
	}

	fmt.Println(data)
	c.JSON(http.StatusOK, gin.H{
		"content": string(data),
	})
}

func Delete(c *gin.Context) {
	filePath := FileDir + c.Param("name")
	err := os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件不存在"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stuatus": "ok",
	})
}
