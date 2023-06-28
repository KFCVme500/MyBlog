package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh_Hans_CN"
	unTrans "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
	"myBlog/utils/errmsg"
	"reflect"
)

func Validate(data any) (string, int) {
	validate := validator.New()
	uni := unTrans.New(zh_Hans_CN.New())

	//翻译方法
	trans, _ := uni.GetTranslator("zh_Hans_CN")
	//注册默认的翻译
	err := zhTrans.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println("err:", err)
	}
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		return label
	})
	//判断是否是结构体
	err = validate.Struct(data)
	if err != nil {
		for _, v := range err.(validator.ValidationErrors) {
			return v.Translate(trans), errmsg.ERROR
		}
	}
	return "", errmsg.SUCCSE
}
