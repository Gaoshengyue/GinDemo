package controllers

import (
	"fmt"
	"gindemo/config"
	"gindemo/db"
	"gindemo/schema"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// UserController 控制器 - 用户
type UserController struct{}

// CreateUser GetUserList 获取用户列表
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var createRequest schema.CreateUserRequest
	// 创建一个自定义的验证器
	v := validator.New()

	// 注册自定义的验证函数
	v.RegisterStructValidation(schema.CheckCreateUserRequest, schema.CreateUserRequest{})
	if err := c.ShouldBind(&createRequest); err != nil {
		// 处理绑定错误

		err := &config.RequestValidationError{Code: 200, Message: err.Error()}
		c.Error(err)
		return
	}
	err := v.Struct(createRequest)
	if err != nil {
		// 处理验证错误
		// 可以根据需要自定义错误处理逻辑
		err := &config.RequestValidationError{Code: 200, Message: err.Error()}
		c.Error(err)
		return
	}
	fmt.Print(createRequest)
	createUserObj := db.MerchantStaff{Username: createRequest.Username, MerchantId: createRequest.MerchantId}
	conn, ok := c.Get("DBConnection")
	if !ok {
		// 错误处理，无法获取数据库连接

		err := &config.DBConnError{Code: 200, Message: "db connection error"}
		c.Error(err)
		//c.AbortWithError(err.Code, err)
		return

	}
	dbConn, ok := conn.(*gorm.DB)
	if !ok {
		// 错误处理，类型断言失败

		err := &config.DBConnError{Code: 200, Message: "db connection error"}
		c.Error(err)
		//c.AbortWithError(err.Code, err)
		return
	}
	err = createUserObj.CreateUser(dbConn)
	fmt.Printf("%v", err)
	if err != nil {
		fmt.Printf("有错误进来了")
		err := &config.UserCreateError{Code: 200, Message: "create user error"}
		c.Error(err)
		//c.AbortWithError(err.Code, err)
		return
	}

	// 模拟从数据库获取用户数据
	//db.MerchantStaff{}

	c.Set("responseData", make(map[string]int))
	return
	//c.JSON(http.StatusOK, createRequest)
}
