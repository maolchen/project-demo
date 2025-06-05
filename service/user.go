package service

import (
	"github.com/gin-gonic/gin"
	"github.com/maolchen/krm-backend/models"
	. "github.com/maolchen/krm-backend/utils"
	"github.com/maolchen/krm-backend/validator"
)

// 创建用户
func CreateUser(ctx *gin.Context) (uint, error) {
	// 接收数据
	user := models.User{}
	var err error
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}
	// 数据校验

	// 校验用户名
	if err := validator.ValidateUsername(user.Username); err != nil {
		return 0, err
	}

	// 校验密码
	if err := validator.ValidatePassword(user.Password); err != nil {
		return 0, err
	}
	// 数据处理
	user.HashPass, err = MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	//写入数据库
	return user.Insert()

}

// 查询用户By ID
func GetUserById(id uint) (user models.User) {
	user = models.User{}
	user.GetOneById(id)
	return user

}

// 查询用户by name
func GetUserByUsername(username string) (user models.User) {
	user = models.User{}
	user.GetOneByUsername(username)
	return user
}

// 查询所有用户
func GetAllUsers() (users []models.User) {
	user := models.User{}
	return user.GetAll()
}

// 更新密码
func ChangeUserPassword(user models.User, RawPassword string) {
	password, _ := MakeHashPassword(RawPassword)
	user.ChangePassword(password)
}
