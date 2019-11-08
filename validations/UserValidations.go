package validations

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"go-cms/models"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
)

type UserValidations struct {
	BaseValidations
}

func init() {

}

func (v *UserValidations) Check(user *models.User) (message string){
	//中文翻译器
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	//验证器
	validate := validator.New()
	//验证器注册翻译器
	zh_translations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			//翻译错误信息
			message = err.Translate(trans)
		}
	}
	return
}