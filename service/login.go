package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/maolchen/project_demo/constants"
	"github.com/maolchen/project_demo/models"
	"github.com/maolchen/project_demo/validator"
	"gorm.io/gorm"
)

// 用户登录
func UserLogin(ctx *gin.Context) (user models.User, err error) {
	// 创建一个接收数据的结构体，不直接使用user，是因为接收了数据ShouldBindJSON后，user中就有数据了，
	// 再通过GetOneByUsername去查数据，即使不存在，gorm也是返回一个空的user，然后这个user又因为被ShouldBindJSON
	// 最后导致user中传过来什么值，就返回什么值
	// GetOneByUsername 改造成返回一个error，查不到就返回error
	var req struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		return user, err
	}

	// 校验用户名
	if err := validator.ValidateUsername(req.Username); err != nil {
		return user, err
	}

	// 验证用户是否存在
	err = user.GetOneByUsername(req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 统一返回“用户名或密码错误”，防止攻击者枚举用户名
			//return user, errors.New(constants.NoSuchUser)
			return user, errors.New(constants.LoginError)
		}
		return user, err
	}

	//// 校验密码
	//if err := validator.ValidatePassword(req.Password); err != nil {
	//	return user, err
	//}

	//验证密码
	ret := user.CheckPassword(req.Username, req.Password)
	if !ret {
		//err = errors.New(constants.PasswordFail)
		// 统一提示用户或密码错误，更安全
		err = errors.New(constants.LoginError)
		return
	}
	return
}
