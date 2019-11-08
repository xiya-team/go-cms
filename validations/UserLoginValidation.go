package validations

type UserLoginValidation struct {
	UserName string `form:"user_name" valid:"Required;"`
	Password string `form:"password" valid:"Required; MinSize(6); MaxSize(16)"`
}
