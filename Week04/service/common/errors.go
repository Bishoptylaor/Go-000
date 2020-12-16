package common

import (
	"errors"
)

var (
	ErrInternalServer   = errors.New("服务异常，请稍后再试")
	ErrParams           = errors.New("参数错误")
	ErrNotSupported     = errors.New("该接口未支持")
	ErrMultiClick       = errors.New("请不要重复点击")
	ErrEventIdMissMatch = errors.New("活动ID匹配失败")
)
