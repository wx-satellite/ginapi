package errno


/*
 错误码为5个数字： 1 表示错误类别（1系统2普通） 00 服务模块代码  01具体的错误代码
参考：https://open.weibo.com/wiki/Error_code
 */

var (
	OK = &Errno{Code: 0, Message: "请求成功！"}
	InternalServerError = &Errno{Code: 10001, Message:"服务器错误！"}
	ErrBind = &Errno{Code: 10002, Message: "结构体参数绑定失败！"}

	ErrValidation = &Errno{Code: 20001, Message:"数据校验失败！"}
	ErrDatabase = &Errno{Code: 20002, Message:"数据库操作失败！"}
	AuthValidation = &Errno{Code:20003, Message:"登录认证失败！"}
	ErrJWTIsNotLegal = &Errno{Code: 20004, Message:"token不合法！"}
	ErrJWTCreateFail = &Errno{Code: 20005, Message:"token生成失败！"}

	// 用户相关
	ErrUserNotFound = &Errno{Code: 20101, Message:"用户不存在！"}
	ErrUserPasswordNotValid = &Errno{Code:20102, Message:"用户名或者密码错误！"}
)
