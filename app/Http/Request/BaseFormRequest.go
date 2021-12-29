package Request

import (
	"fmt"
	"order/app/Common"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

var Trans ut.Translator

func RemoveTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {

		zhTrans := zh.New()                      // 中文转换器
		enTrans := en.New()                      // 因为转换器
		uni := ut.New(zhTrans, zhTrans, enTrans) // 创建一个通用转换器
		curLocales := locale                     // 设置当前语言类型
		var ok bool
		Trans, ok = uni.GetTranslator(curLocales) // 获取对应语言的转换器
		// validate := validator.New()
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
			if name == "-" {
				return ""
			}

			return name
		})
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// 在校验器注册自定义的校验方法
		if err := v.RegisterValidation("check_date", checkDate); err != nil {
			return err
		}
		switch curLocales {
		case "zh":
			// 内置tag注册 中文翻译器
			_ = zh_trans.RegisterDefaultTranslations(v, Trans)
			// 自定义tag注册 中文翻译器
			_ = v.RegisterTranslation("unique_name", Trans, registerTranslator("unique_name", "{0}已经被占用"), translate)
			_ = v.RegisterTranslation("check_date", Trans, registerTranslator("check_date", "{0}必须要晚于当前日期"), translate)
		case "en":
			// 内置tag注册 英文翻译器
			_ = en_trans.RegisterDefaultTranslations(v, Trans)
		}
	}
	return
}

func checkDate(fl validator.FieldLevel) bool {
	format := "2006-01-02"
	now := time.Now()

	date, err := time.Parse(format, fl.Field().String())
	// color.Debug.Println(date)
	b := now.Format(format)
	a := date.Format(format)
	// color.Debug.Println(b)
	if err != nil {
		return false
	}
	// color.Debug.Println(a)
	// color.Debug.Println(b)
	s1 := Common.DateToTimeS(a, "Y-m-d")
	s2 := Common.DateToTimeS(b, "Y-m-d")

	if s1 > s2 {
		// color.Danger.Println(time.Now())
		return false
	}
	return true
}

// registerTranslator 为自定义字段添加翻译功能
func registerTranslator(tag string, msg string) validator.RegisterTranslationsFunc {
	return func(trans ut.Translator) error {
		if err := trans.Add(tag, msg, false); err != nil {
			return err
		}
		return nil
	}
}

// translate 自定义字段的翻译方法
func translate(trans ut.Translator, fe validator.FieldError) string {
	msg, err := trans.T(fe.Tag(), fe.Field())
	if err != nil {
		panic(fe.(error).Error())
	}
	return msg
}
