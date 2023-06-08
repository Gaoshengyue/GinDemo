package middlewares

import (
	"gindemo/db"
	"github.com/gin-gonic/gin"
)

func GetDBConnection(c *gin.Context) {
	// 从连接池获取连接
	conn := db.GetDBConnection()

	// 将连接绑定到上下文中
	c.Set("DBConnection", conn)

	// 继续处理请求
	c.Next()

	// 在请求结束后释放连接
	db.ReleaseDBConnection(conn)
}
