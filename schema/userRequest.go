package schema

import (
	"github.com/go-playground/validator/v10"
)

type CreateUserRequest struct {
	Username string `form:"username" binding:"required"`
	//Username   string `form:"username" binding:"required,min=5"`  //两种写法
	MerchantId int `form:"merchant_id" binding:"required"`
}

func CheckCreateUserRequest(sl validator.StructLevel) {
	req := sl.Current().Interface().(CreateUserRequest)
	// 进行自定义校验逻辑
	if len(req.Username) < 5 || len(req.Username) > 20 {
		sl.ReportError(req.Username, "Username", "Username", "length必须在5到20之间", "")
	}
}
