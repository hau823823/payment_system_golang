package dao

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	 _ "github.com/go-sql-driver/mysql"
  )

//指定驅動
const DRIVER = "mysql"

var SqlSession *gorm.DB

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Missing .env file.")
	}
}

//初始化連線資料庫，生成可操作基本增刪改查結構的變數
func InitMySql()(err error)  {

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("USER_NAME"),
		os.Getenv("USER_PASSWORD"),
		os.Getenv("URL"),
		os.Getenv("PORT"),
		os.Getenv("DB_NAME"),
	)
	//連線資料庫
	SqlSession,err =gorm.Open(DRIVER,dsn)
	if err !=nil{
		panic(err)
	}
	//驗證資料庫連線是否成功，若成功，則無異常
	return SqlSession.DB().Ping()
}

//關閉資料庫連線
func Close()  {
	SqlSession.Close()
}