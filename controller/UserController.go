package controller

import (
	"time"
	"log"
	"os"
	"cash_register_system/entity"
	"cash_register_system/model"
	"cash_register_system/common"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)

//初始化 .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Missing .env file.")
	}
}

//實作登入功能
func LoginUserByEmailAndPassword(c *gin.Context) {
	var user entity.User

	c.BindJSON(&user)
	err := model.LoginUserByEmailAndPassword(&user)
	if err!=nil {
		rsp.Error(c,err.Error())
	}else {
		//rsp.Success(c,"登入成功",user)
		//JWT
		token, err := GenerateToken(user.Id, user.Name)
		c.Writer.Header().Set("Token", token)
		c.Next()
		if err!=nil {
			rsp.Error(c,err.Error())
		}else {
			rsp.Success(c,"登入成功",token)
		}
	}
}

//新增user資訊
func CreateUser(c *gin.Context)  {
	var user entity.User

	c.BindJSON(&user)
	err := model.CreateUser(&user)
	if err!=nil {
		rsp.Error(c,err.Error())
	}else {
		rsp.Success(c,"新增成功",user)
	}
}

//通過jwt傳輸的數據，包含UserId 和 Username
type Claims struct {
	Id     int64
	Name   string
	jwt.StandardClaims
}

//生成token
func GenerateToken(uid int64, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(2 * time.Hour)
	issuer := "HAU"
	claims := Claims{
	  Id:    uid,
	  Name:  username,
	  StandardClaims: jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		Issuer:    issuer,
	  },
	}
  
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(os.Getenv("MY_SECRET")))
	return token, err
}

//解析token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
	  return []byte(os.Getenv("MY_SECRET")), nil
	})
	if err != nil {
	  return nil, err
	}
  
	if tokenClaims != nil {
	  if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		return claims, nil
	  }
	}
  
	return nil, err
}

//透過解析JWT獲得Id，進而獲得其他資訊
func GetUserByToken(c *gin.Context, token string) (user entity.User, err error) {
	u_claims, err_token := ParseToken(token)
	if err_token!=nil {
		rsp.Error(c,err_token.Error())
	}

	user,err = model.GetUserById(u_claims.Id)
	return
}