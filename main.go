package main

import (
	"nirs/packages/forecast"

	"github.com/gin-gonic/gin"
)

const (
	Port = ":1010"
)

var router *gin.Engine

type outData struct {
	Name    string
	Percent int
	Color   int
}

type inData struct {
	Region   string `json:"Region"`
	District string `json:"District"`
}

var odata outData
var idata inData

func main() {
	router = gin.Default()
	router.Static("/assets/", "assets/")
	router.Static("/src/", "src/")
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/", handlerIndex)
	router.POST("/calculation", handlerCalculation)
	_ = router.Run(Port)
}

func handlerIndex(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"Region":   idata.Region,
		"District": odata.Name,
		"Percent":  odata.Percent,
		"Color":    odata.Color,
	})
}

func handlerCalculation(c *gin.Context) {
	err := c.BindJSON(&idata)
	if err != nil {
		c.JSON(400, gin.H{})
		return
	}
	odata.Name, odata.Percent, odata.Color = forecast.CheckAccidentRate(idata.Region, idata.District)
}
