package Example

// 拦截器

import (
	"log"
	"net/http"
	"order/app/Common"
	"order/app/Http/Middlewares"
	"order/app/Http/Models/Kit"
	"order/app/Http/Models/User"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"github.com/gookit/color"
)

func VerifyExample(ctx *gin.Context) {
	ctx.Next()
}

func Test(ctx *gin.Context) {
	method := ctx.Request.Method
	body := ctx.Request.Body
	header := ctx.Request.Header["Content-Type"]
	//
	//fmt.Println(ctx.Request.Form)
	//
	//fmt.Println("====3===")

	id := "123" //
	nickname := "456"
	// 接口返回
	back := gin.H{
		"method":   method,
		"body":     body,
		"header":   header,
		"id":       id,
		"nickname": nickname,
		//"latency": ctx.Get("state_latency"),
		//"img": img,
		//"excel_name": excelName,
		//"excel_state": excelState,
		////"timezone": Common.ServerInfo["timezone"],
		//"date": Common.GetTimeDate("Y-m-d H:i:s"),

	}
	ctx.JSON(200, back)
}

type UserInfo struct {
	Uname  string `validate:"required,min=5" label:"用户名"`       // UUID 类型
	Passwd string `validate:"required,min=6,max=32" label:"密码"` // 自定义校验

}

func CheckUniqueName(fl validator.FieldLevel) bool {
	// 获取字段当前值 fl.Field()
	// 获取tag 对应的参数 fl.Param() ，针对unique_name标签 ，不需要参数
	// 获取字段名称 fl.FieldName()

	// balabala处理一波，比如查库比较

	return false
}

// Login 用户登录
// Usage:
// 	msg := Login(*gin.Context)
// 	@auther yeweipeng 2021-07-27
// 	@params *gin.Context
// 	@return  null
func Login(ctx *gin.Context) {
	uname := Common.Input(ctx, "uname")
	passwd := Common.Input(ctx, "passwd")
	autoLogin := Common.Input(ctx, "auto_login")
	autoFlag, _ := strconv.ParseBool(autoLogin)
	s := UserInfo{
		Uname:  uname,
		Passwd: passwd,
	}
	color.Info.Printf("%T =>%v", s, s)
	// 创建翻译器
	zhTrans := zh.New()                       // 中文转换器
	enTrans := en.New()                       // 因为转换器
	uni := ut.New(zhTrans, zhTrans, enTrans)  // 创建一个通用转换器
	curLocales := "zh"                        // 设置当前语言类型
	trans, _ := uni.GetTranslator(curLocales) // 获取对应语言的转换器

	validate := validator.New()                                     // 创建验证器
	_ = validate.RegisterValidation("unique_name", CheckUniqueName) // 注册自定义tag回调函数

	switch curLocales {
	case "zh":
		// 内置tag注册 中文翻译器
		_ = zh_trans.RegisterDefaultTranslations(validate, trans)
		// 自定义tag注册 中文翻译器
		_ = validate.RegisterTranslation("unique_name", trans, func(ut ut.Translator) error {
			if err := ut.Add("unique_name", "{0}已被占用", false); err != nil {
				return err
			}
			return nil
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, err := ut.T(fe.Tag(), fe.Field())
			if err != nil {
				log.Printf("警告: 翻译字段错误: %#v", fe)
				return fe.(error).Error()
			}
			return t
		})
	case "en":
		// 内置tag注册 英文翻译器
		_ = en_trans.RegisterDefaultTranslations(validate, trans)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}

		return name
	})
	// validate := validator.New()
	err := validate.Struct(s) // 执行验证
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			// color.Error.Println(e.Translate(trans))
			ctx.AbortWithStatusJSON(200, gin.H{
				"code":    500,
				"token":   "",
				"message": e.Translate(trans),
			})
			break
		}
		return
	}

	resultInfo, err := User.GetUserInfo(uname)
	// color.Debug.Println(validate)
	color.Debug.Println(resultInfo)
	color.Debug.Println(resultInfo.Mobile)
	if resultInfo != nil && err == nil {
		if uname == resultInfo.Mobile && passwd == Common.MD5("123456") {
			m := new(Middlewares.User)
			m.Name = resultInfo.StaffName
			m.Avatar = resultInfo.Avatar
			m.Id = resultInfo.StaffId
			m.Mobile = resultInfo.Mobile
			var d time.Duration
			if autoFlag {
				d = time.Second * 5 * 24 * 3600
			} else {
				d = time.Second * 1800
			}
			token, err := Middlewares.JwtGenerateToken(m, d)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
			}
			Kit.RDB.Set(ctx, "go_database_jwt_cache:"+uname, token, d)
			ctx.JSON(200, gin.H{
				"code":    200,
				"token":   token,
				"message": "success",
			})
		} else {
			ctx.JSON(200, gin.H{
				"code":    500,
				"token":   "",
				"message": "帐号密码错误",
			})
		}
	} else if resultInfo == nil && err == nil {
		ctx.JSON(200, gin.H{
			"code":    404,
			"token":   "",
			"message": "用户不存在",
		})
	} else if resultInfo == nil && err != nil {
		ctx.JSON(200, gin.H{
			"code":    500,
			"token":   "",
			"message": err,
		})
	}
	// 接口返回

	// fmt.Printf("时间v1 %+v", time.Hour*24*365)
	// fmt.Printf("时间v2 %+v", time.Second*180)

	// token, err := jwtGenerateToken(m, time.Hour*24*365)
	// token, err := jwtGenerateToken(m, time.Second*30)
	// if err != nil {
	// 	ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
	// }
	// ctx.Header("JWT", token)
}

type MsgJson struct {
	Token string `json:"token"`
}

func RefreshToken(ctx *gin.Context) {
	// token := Common.Input(ctx, "token")
	inputJson := MsgJson{}
	if err := ctx.ShouldBindJSON(&inputJson); err != nil {
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
		return
	}
	user, err := Middlewares.JwtParseExpireUser(inputJson.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
		return
	} else {
		var d time.Duration

		uname := user.Mobile
		d = time.Second * 1800

		newToken, err := Middlewares.JwtGenerateToken(user, d)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"msg": err.Error()})
			return
		}
		Kit.RDB.Set(ctx, "go_database_jwt_cache:"+uname, newToken, d)
		// 接口返回
		back := gin.H{
			"token": newToken,
		}
		ctx.JSON(200, back)
	}
}
