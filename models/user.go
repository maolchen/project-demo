package models

import (
	"github.com/maolchen/project_demo/controllers/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;comment:'用户名'" json:"username,omitempty"`
	HashPass string `gorm:"not null;size:255;comment:'密码'" json:"-"`
	Password string `gorm:"-" json:"password,omitempty"`
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
func (user *User) GetOneByUsername(username string) {
	db := database.GetDB()
	db.Where("username = ?", username).First(&user)
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
