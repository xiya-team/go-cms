package resultError

import (
	"go-cms/resultModels"
)

type IError interface {
	Error() string
	GetErrCode() resultModels.ErrorCode
}

// 错误
type FundingError struct {
	Code resultModels.ErrorCode
	Msg  string
}

func NewFallFundingErr(msg string) *FundingError {
	return &FundingError{Code: resultModels.FALL, Msg: msg}
}

func (e *FundingError) Error() string {
	return e.Msg
}

func (e *FundingError) GetErrCode() resultModels.ErrorCode {
	return e.Code
}

// 表单错误
var (
	FormParamErr = FundingError{Code: resultModels.FALL, Msg: "请求参数有误"}
)

// 登录和身份验证相关错误
var (
	NotLoginError       = FundingError{Code: resultModels.NOT_LOGIN, Msg: "没有登录"}
	UserNotExitError    = FundingError{Code: resultModels.FALL, Msg: "用户不存在"}
	UserPasswordError   = FundingError{Code: resultModels.FALL, Msg: "账号或密码错误"}
	UserRoleVerifyError = FundingError{Code: resultModels.VERIFY_FALL, Msg: "用户身份验证失败"}
)

// 地址相关
var (
	AddressNotFound = FundingError{Code: resultModels.FALL, Msg: "没有找到对应地址"}
	AddressInfoErr  = FundingError{Code: resultModels.FALL, Msg: "地址信息有误"}
)

// 产品相关错误
var (
	ProductNotFound    = FundingError{Code: resultModels.FALL, Msg: "没有找到相关产品"}
	ProductPkgNotFound = FundingError{Code: resultModels.FALL, Msg: "没有找到相关套餐"}
	OutOfStock         = FundingError{Code: resultModels.FALL, Msg: "库存不足"}
)

// 订单相关错误
var (
	OrderCreateErr = FundingError{Code: resultModels.FALL, Msg: "创建订单时发生错误"}
)
