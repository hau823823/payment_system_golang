package main

import (
	"cash_register_system/dao"
	"cash_register_system/entity"
	"cash_register_system/routers"
	_ "github.com/go-sql-driver/mysql"
)

func main()  {
	//連線資料庫
	err := dao.InitMySql()
	if err !=nil{
	   panic(err)
	}
	//程式退出關閉資料庫連線
	defer dao.Close()
	//綁定模型
	dao.SqlSession.AutoMigrate(&entity.User{}, &entity.Payment_coin{})
	//註冊路由
	router := routers.SetRouter()
	//啟動埠為6666的專案
	router.Run(":6666")
 }