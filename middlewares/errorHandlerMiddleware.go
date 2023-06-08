package middlewares

import (
	"fmt"
	"gindemo/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinGlobalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 在处理请求之前执行的逻辑

		c.Next()
		// 在处理完请求后检查是否有错误发生
		err := c.Errors.Last()
		if err != nil {
			fmt.Print(err)
			// 判断是否为自定义错误类型
			switch customErr := err.Err.(type) {
			case *config.UserCreateError:
				// 处理自定义错误
				c.JSON(customErr.Code, gin.H{"message": customErr.Message, "code": 1})
				return
			case *config.RequestValidationError:
				c.JSON(customErr.Code, gin.H{"message": customErr.Message, "code": 1})
				return

			}

			// 根据错误类型返回不同的错误响应
			switch err.Type {
			case gin.ErrorTypeBind:
				// 处理绑定错误
				c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数有误", "code": 1})
			case gin.ErrorTypeRender:
				// 处理渲染错误
				c.JSON(http.StatusInternalServerError, gin.H{"message": "服务器内部错误", "code": 1})

			default:
				// 其他类型的错误处理
				c.JSON(http.StatusInternalServerError, gin.H{"message": "未知错误", "code": 1})
			}
		} else {
			// 对返回结果进行统一处理
			// 例如添加通用的响应头、格式化响应数据等
			// ...

			// 示例：添加通用的响应头
			c.Header("X-Custom-Header", "value")

			// 示例：格式化响应数据
			data := c.MustGet("responseData")

			c.JSON(http.StatusOK, gin.H{"code": 0, "message": "请求成功", "data": data})
		}

	}
}
