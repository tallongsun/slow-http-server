package main

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		var ts uint64
		tsStr, ok := c.GetQuery("ts")
		if ok {
			m, err := strconv.ParseUint(tsStr, 10, 64)
			if err != nil {
				ts = 100
			}
			ts = m
		}

		result, err := sleep(c.Request.Context(), ts)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, result)
	})
	app.Run(":80")
}
func sleep(ctx context.Context, max uint64) (string,error){
	logrus.Infof("Start sleep ts: %d", max)
	time.Sleep(time.Duration(max) * time.Millisecond)
	return "ok",nil
}

