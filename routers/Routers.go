package routers

import (
	"cash_register_system/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	//user路由組
	userGroup := r.Group("user")
	{
	   //註冊User
	   userGroup.POST("/regist",controller.CreateUser)
	   //登入User
	   userGroup.POST("/login",controller.LoginUserByEmailAndPassword)
	}

	//payment路由組
	paymentGroup := r.Group("payment")
	{
	   //平台幣付款
	   paymentGroup.POST("/coin",controller.PayPlatformCoin)
	   //平台點數付款
	   paymentGroup.POST("/point",controller.PayPlatformPoint)
	}

	return r
}