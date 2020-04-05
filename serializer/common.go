package serializer

import "github.com/gin-gonic/gin"

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// TrackedErrorResponse 有追踪信息的错误响应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
// 5开头的五位数错误编码为服务器端错误，比如数据库操作失败
// 4开头的五位数错误编码为客户端错误，有时候是客户端代码写错了，有时候是用户操作错误
const (
	// CodeNoRightErr 未授权访问
	CodeNoRightErr = 403

	//CodeRegisterErr 用户注册信息错误
	CodeRegisterErr = 40001
	// CodeNeverLogin 用户未登录
	CodeNeverLogin = 40002
	// CodeInvalidToken 无效的Token,重新登录获取Token
	CodeInvalidToken = 40003
	//CodeParamErr 各种奇奇怪怪的参数错误
	CodeParamErr = 40004
	// CodeDBError 数据库操作失败,联系管理员
	CodeDBError = 50001
	// CodeEncryptError 加密失败,联系管理员
	CodeEncryptError = 50002
	// CodeLoseAuthHeader 请求头中缺少Auth Token，联系前端管理员
	CodeLoseAuthHeader = 50003
	// CodeAuthWrong 请求头中Auth Token格式有误，联系前端管理员
	CodeAuthWrong = 50004
	// CodeUnCreateToken 生成Token失败,联系后端管理员
	CodeUnCreateToken = 50005
)

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Status: errCode,
		Msg:    msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && gin.Mode() != gin.ReleaseMode {
		res.Error = err.Error()
	}
	return res
}

// DBErr 数据库操作失败
func DBErr(msg string, err error) Response {
	if msg == "" {
		msg = "数据库操作失败"
	}
	return Err(CodeDBError, msg, err)
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}
