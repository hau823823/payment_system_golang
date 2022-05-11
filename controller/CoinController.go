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

//平台幣付款
func PayPlatformCoin(c *gin.Context)  {
	//JWT驗證
	token := c.Request.Header.Get("Token")
	user, err_token := GetUserByToken(c, token)
	if err_token!=nil {
		rsp.Error(c,err_token.Error())
	}
	
	//request body：輸入商品價格
	var input struct {
		Price float64 `json:"price"`
	}
	c.BindJSON(&input)

	//取得使用者當下平台幣帳戶資訊
	account_info,err_account := model.GetCoinAccountInfoByUid(user.Id)
	if err_account!=nil {
		rsp.Error(c,err_account.Error())
	}

	//判斷是否為VIP
	//非vip
	if user.Vip == 0 {
		//寫入資料庫的資料格式
		payment_coin := entity.Payment_coin {
			Uid      :  user.Id,
			Account  :  account_info.Account - input.Price,
			Cost     :  input.Price,
		}

		if payment_coin.Account >= 0 {
			err := model.PayPlatformCoin(&payment_coin)
			if err!=nil {
				rsp.Error(c,err.Error())
			}else {
				rsp.Success(c,"一般付款成功",payment_coin)
			}
		} else {
			rsp.Error(c,"平台幣餘額不足")
		}
	} else {
		//VIP
		//取得折扣資訊
		discount_info,err_discount := model.GetDiscountCoinVIP(user.Vip)
		if err_account!=nil {
			rsp.Error(c,err_discount.Error())
		}
		
		//寫入資料庫的資料格式
		payment_coin := entity.Payment_coin {
			Uid      :  user.Id,
			Account  :  account_info.Account - input.Price * discount_info.Disacount,
			Cost     :  input.Price * discount_info.Disacount,
		}

		if payment_coin.Account >= 0 {
			err := model.PayPlatformCoin(&payment_coin)
			if err!=nil {
				rsp.Error(c,err.Error())
			}else {
				rsp.Success(c,"折扣付款成功",payment_coin)
			}
		} else {
			rsp.Error(c,"平台幣餘額不足")
		}
	}
}
