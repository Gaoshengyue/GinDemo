package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

func LoggerMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 请求开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 请求结束时间
		endTime := time.Now()
		// 读取请求体的内容
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// 处理读取请求体的错误
		}
		bodyString := string(body)
		// 记录请求的详细信息
		logger.Infof("[GIN] %v | %3d | %13v | %15s | %-7s %s | %+v\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			c.Writer.Status(),
			endTime.Sub(startTime).String(),
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			bodyString,
		)
	}
}
