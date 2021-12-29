package Middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	AppSecret = ""                          // viper.GetString会设置这个值(32byte长度)
	AppIss    = "github.com/libragen/felix" // 这个值会被viper.GetString重写
)

type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
}

type userStdClaims struct {
	jwt.StandardClaims
	*User
}

func CheckLogin(ctx *gin.Context) {
	// fmt.Printf("时间v1 %+v", time.Hour*24*365)
	// fmt.Printf("时间v2 %+v", time.Second*180)

	// token, err := jwtGenerateToken(m, time.Hour*24*365)
	// token, err := jwtGenerateToken(m, time.Second*30)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
	// }
	// ctx.Header("JWT", token)
	GetCheckJwtToken(ctx)
	ctx.Next()
}

const (
	contextKeyUserObj = "authedUserObj"
	bearerLength      = len("Bearer ")
)

func GetCheckJwtToken(ctx *gin.Context) {
	hToken := ctx.GetHeader("Authorization")
	if len(hToken) < bearerLength {
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": "Authorization header  has not token"})
		return
	}

	// token := strings.TrimSpace(hToken[bearerLength:])
	// 解析token
	fmt.Println(hToken)
	usr, err := JwtParseUser(hToken)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
		return
	}
	ctx.Set(contextKeyUserObj, *usr)
	ctx.Next()
}

func UserLogin(ctx *gin.Context) {
	ctx.Next()
}

func JwtGenerateToken(m *User, d time.Duration) (string, error) {
	m.Password = ""
	// m.Id = 1000
	// m.Name = "wismyzhizi"
	// m.Age = 18
	// m.Password = ""
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%v", m.Id),
		Issuer:    AppIss,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           m,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		// log.WithError(err).Fatal("config is wrong, can not generate jwt")
		fmt.Println("config is wrong, can not generate jwt")
	}
	return tokenString, err
}

// JwtParseUser 解析未过期的token,如果过期直接报错提示
func JwtParseUser(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := userStdClaims{}
	// fmt.Println("开始解析token" + tokenString)
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	// log.Println(err)
	if err != nil {
		return nil, err
	}
	// log.Printf("结束解析token%+v", claims.StandardClaims)
	return claims.User, err
}

// JwtParseExpireUser 解析过期的token
func JwtParseExpireUser(tokenString string) (*User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found")
	}
	claims := userStdClaims{}
	// fmt.Println("开始解析token" + tokenString)
	_, _, err := new(jwt.Parser).ParseUnverified(tokenString, &claims)
	// log.Println(err)
	if err != nil {
		return nil, err
	}
	// log.Printf("结束解析token%+v", claims.StandardClaims)
	return claims.User, err
}
