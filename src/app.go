package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/api/bixi-around", func(c *gin.Context) {
		city := c.DefaultQuery("city", "montreal")
		nbResult, _ := strconv.ParseInt(c.DefaultQuery("nbResult", "5"), 0, 16)
		lat, _ := strconv.ParseFloat(c.DefaultQuery("lat", "45.501690"), 64)
		lon, _ := strconv.ParseFloat(c.DefaultQuery("lon", "-73.567253"), 64)
		c.JSON(http.StatusOK, bixiLogic(city, int(nbResult), lat, lon))
	})
	router.Run()
}
