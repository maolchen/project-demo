package models

import (
	"errors"
	"github.com/maolchen/krm-backend/database"
)

func UpdateByName(model interface{}, name string, updates map[string]interface{}) error {
	return database.GetDB().
		Model(model).
		Where("name = ?", name).
		Updates(updates).Error
}

func DeleteByName(model interface{}, name string) error {
	db := database.GetDB().Unscoped().Where("name = ?", name).Delete(model)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		return errors.New("集群不存在或已被删除")
	}
	return nil
}
