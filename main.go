package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"wiki-service/model"
)

func main() {
	router := gin.Default()

	router.GET("/ht", func(c *gin.Context) {
		eventType, _ := strconv.Atoi(c.DefaultQuery("type", "-1"))
		month, _ := strconv.Atoi(c.DefaultQuery("month", "1"))
		day, _ := strconv.Atoi(c.DefaultQuery("day", "1"))
		index, _ := strconv.Atoi(c.Query("pageIndex"))
		size, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
		rsp, err := model.FetchEvents(eventType, month, day, index, size)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"msg": "error",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"msg": "",
				"content": gin.H{
					"date": strconv.Itoa(month) + "-" + strconv.Itoa(day),
					"events": rsp,
				},
			})
		}
	})

	router.Run()
}
