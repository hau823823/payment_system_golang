package controller

import (
	"log"
	"cash_register_system/entity"
	"cash_register_system/model"
	"cash_register_system/common"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//初始化 .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Missing .env file.")
	}
}

//平台點數付款
func PayPlatformPoint(c *gin.Context)  {
	//JWT驗證
	token := c.Request.Header.Get("Token")
	user, err_token := GetUserByToken(c, token)
	if err_token!=nil {
		rsp.Error(c,err_token.Error())
	}
	
	//request body：輸入商品價格
	var input struct {
		Price float64 `json:"price"`
		Point float64 `json:"point"`
	}
	c.BindJSON(&input)

	//取得使用者當下平台點數帳戶資訊
	point_account_info,err_account1 := model.GetPointAccountInfoByUid(user.Id)
	if err_account1!=nil {
		rsp.Error(c,err_account1.Error())
	}

	//取得使用者當下平台幣帳戶資訊
	coin_account_info, err_account2 := model.GetCoinAccountInfoByUid(user.Id)
	if err_account2!=nil {
		rsp.Error(c,err_account2.Error())
	}

	//取得使用平台點數折抵平台幣比率
	point_discount, err_percentage := model.GetDiscountPoint()
	if err_percentage!=nil {
		rsp.Error(c,err_percentage.Error())
	}

	if input.Point >= 100  && user.Vip > 0 {
		payment_point := entity.Payment_point {
			Uid      :  user.Id,
			Account  :  point_account_info.Account - input.Point,
			Cost     :  input.Point,
		}
		payment_coin := entity.Payment_coin {
			Uid      :  user.Id,
			Account  :  coin_account_info.Account - (input.Price - (input.Point * point_discount.Percentage)) * 0.9,
			Cost     :  (input.Price - (input.Point * point_discount.Percentage)) * 0.9,
		}
	
		if payment_point.Account >= 0 && payment_coin.Account >= 0{
			err := model.PayPlatformPoint(&payment_point)
			if err!=nil {
				rsp.Error(c,err.Error())
			}else {
				rsp.Success(c,"點數扣款成功",payment_point)
			}
	
			err = model.PayPlatformCoin(&payment_coin)
			if err!=nil {
				rsp.Error(c,err.Error())
			}else {
				rsp.Success(c,"平台幣扣款成功",payment_coin)
			}
		} else {
			rsp.Error(c,"平台幣或平台點數餘額不足")
		} 
	} else {
		payment_point := entity.Payment_point {
			Uid      :  user.Id,
			Account  :  point_account_info.Account - input.Point,
			Cost     :  input.Point,
		}
		payment_coin := entity.Payment_coin {
			Uid      :  user.Id,
			Account  :  coin_account_info.Account - input.Price + (input.Point * point_discount.Percentage),
			Cost     :  input.Price - (input.Point * point_discount.Percentage),
		}
	
		if payment_point.Account >= 0 && payment_coin.Account >= 0{
			err := model.PayPlatformPoint(&payment_point)
			if err!=nil {
				rsp.Error(c,err.Error())
			}else {
				rsp.Success(c,"點數扣款成功",payment_point)
			}
	
			err = model.PayPlatformCoin(&payment_coin)
			if err!=nil {
				rsp.Error(c,err.Error())
			}else {
				rsp.Success(c,"平台幣扣款成功",payment_coin)
			}
		} else {
			rsp.Error(c,"平台幣或平台點數餘額不足")
		} 
	}
}
