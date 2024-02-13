package errors

import (
	"libong/common/server/http/code"
)

var (
	ParamsError    = code.Error(90000, "参数为空或错误")
	ImgFormatError = code.Error(90004, "图片格式错误")
)
