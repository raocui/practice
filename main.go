package main

import (
	"_/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/maybgit/glog"
	"net/http"
	"time"
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
	glog.Error("aaaa")
	sh, _ := time.LoadLocation("Asia/Shanghai")
	time.Local = sh
	r := gin.Default()
	mygorm := api.Mygorm{}
	r.GET("/gorm/migration", mygorm.Migration)
	r.POST("/gorm/create", mygorm.Create)

	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println(json)
		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
	})
	r.Run()
}
