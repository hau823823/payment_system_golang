package model

import (
	"cash_register_system/dao"
	"cash_register_system/entity"
)

//新增平台幣付款資訊
func PayPlatformCoin(payment *entity.Payment_coin) (err error) {
	if err = dao.SqlSession.Create(payment).Error; err != nil {
		return err
	}
	return
}

//新增平台點數付款資訊
func PayPlatformPoint(payment *entity.Payment_point) (err error) {
	if err = dao.SqlSession.Create(payment).Error; err != nil {
		return err
	}
	return
}

//根據使用者id查詢平台幣帳戶資訊
func GetCoinAccountInfoByUid(uid int64) (payment_info entity.Payment_coin, err error) {
	if err = dao.SqlSession.Where("uid = ?", uid).Last(&payment_info).Error; err != nil {
		return payment_info, err
	}
	return
}

//根據使用者id查詢平台點數帳戶資訊
func GetPointAccountInfoByUid(uid int64) (payment_info entity.Payment_point, err error) {
	if err = dao.SqlSession.Where("uid = ?", uid).Last(&payment_info).Error; err != nil {
		return payment_info, err
	}
	return
}