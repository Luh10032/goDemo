package filserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Read(c *gin.Context) {
	var body readBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := ioutil.ReadFile(body.File)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件不存在"})
		return
	}

	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{
		"content": string(data),
	})
}

func Create(c *gin.Context) {
	var body createBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := os.Create(body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "文件创建失败"})
		return
	}
	defer file.Close()

	_, err = file.WriteString(body.Content)

	fmt.Println(body)
	c.JSON(http.StatusOK, gin.H{
		"stuatus": "ok",
	})
}
