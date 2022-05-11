package model

import (
	"cash_register_system/dao"
	"cash_register_system/entity"
)

//新增user資訊
func CreateUser(user *entity.User) (err error) {
	if err = dao.SqlSession.Create(user).Error; err != nil {
		return err
	}
	return
}

//登入user
func LoginUserByEmailAndPassword(user *entity.User) (err error) {
	if err = dao.SqlSession.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return err
	}
	return
}

//根據id查詢user
func GetUserById(id int64) (user entity.User, err error) {
	if err = dao.SqlSession.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return
}