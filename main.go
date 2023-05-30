package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var DistrictName string
var DistrictPercent int
var data Data

type Data struct {
	Region   string `json:"Region"`
	District string `json:"District"`
}

func main() {
	router = gin.Default()
	router.Static("/assets/", "assets/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", handlerIndex)
	router.POST("/calculation", handlerCalculation)
	_ = router.Run(":4040")
}

func handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Region":   data.Region,
		"District": DistrictName,
		"Percent":  DistrictPercent,
	})
}

func handlerCalculation(c *gin.Context) {
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}
	DistrictName, DistrictPercent = CheckAccidentRate(data.Region, data.District)
}
