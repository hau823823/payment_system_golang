package model

import (
	"cash_register_system/dao"
	"cash_register_system/entity"
)

//根據VIP查詢平台幣付款折扣比率
func GetDiscountCoinVIP(vip int64) (discount_info entity.Discount_vip_coin, err error) {
	if err = dao.SqlSession.Where("vip = ?", vip).Last(&discount_info).Error; err != nil {
		return discount_info, err
	}
	return
}

//取得平台點數折抵平台幣比率
func GetDiscountPoint() (discount_info entity.Percentage_point, err error) {
	if err = dao.SqlSession.First(&discount_info).Error;err!=nil{
		return discount_info, err
	}
	return
}

//取得VIP使用平台點數後
func GetDiscountPointVIP() (discount_info entity.Discount_vip_point, err error) {
	if err = dao.SqlSession.First(&discount_info).Error;err!=nil{
		return discount_info, err
	}
	return
}