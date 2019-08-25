package controllers

// ErrorController 定义一个通用的错误处理控制器
type ErrorController struct {
	CommonController
}

// Error401 设置html 401错误页面
func (api *ErrorController) Error401() {
	api.SetTpl("error/401.tpl")
}

// Error403 设置html 403错误页面
func (api *ErrorController) Error403() {
	api.SetTpl("error/403.tpl")
}

// Error404 设置html 404错误页面
func (api *ErrorController) Error404() {
	api.SetTpl("error/404.tpl")
}

// ErrorAjax401 设置ajax 401错误信息
func (api *ErrorController) ErrorAjax401() {
	api.TplName = "error/ajax_401.tpl"
}

// ErrorAjax403 设置ajax 403错误信息
func (api *ErrorController) ErrorAjax403() {
	api.TplName = "error/ajax_403.tpl"
}

// ErrorAjax404 设置ajax 404错误信息
func (api *ErrorController) ErrorAjax404() {
	api.TplName = "error/ajax_404.tpl"
}
