package models

import (
	"github.com/maolchen/krm-backend/database"
	"github.com/maolchen/krm-backend/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;comment:'用户名'" json:"username" validate:"required,username"`
	HashPass string `gorm:"not null;size:255;comment:'密码'" json:"-"`
	Password string `gorm:"-" json:"password,omitempty" validate:"required,password"` //omitempty 忽略为空
}

// 设置数据库表名
func (user *User) TableName() string {
	return "user"
}

// 创建用户
func (user *User) Insert() (id uint, err error) {
	db := database.GetDB()
	res := db.Create(&user)
	return user.ID, res.Error

}

// 根据指定ID获取用户
func (user *User) GetOneById(id uint) {
	db := database.GetDB()
	db.First(&user, id)
}

// 根据用户名获取用户信息
func (user *User) GetOneByUsername(username string) error {
	db := database.GetDB()
	return db.Where("username = ?", username).First(user).Error
}

// 获取所有用户信息
func (user *User) GetAll() []User {
	db := database.GetDB()
	var users []User
	db.Find(&users)
	return users
}

// 修改用户密码
func (user *User) ChangePassword(newHashPass string) {
	db := database.GetDB()
	user.HashPass = newHashPass
	db.Save(&user)
}

func (user *User) CheckPassword(username string, password string) bool {
	db := database.GetDB()
	var dbUser User
	if err := db.Where("username = ?", username).First(&dbUser).Error; err != nil {
		return false
	}

	// 使用 bcrypt 直接比对
	return utils.CompareHashAndPassword(dbUser.HashPass, password)
}
