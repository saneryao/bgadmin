package controllers

// 定义一个
type ErrorController struct {
	CommonController
}

func (this *ErrorController) Error401() {
	this.SetTpl("error/401.tpl")
}

func (this *ErrorController) Error403() {
	this.SetTpl("error/403.tpl")
}

func (this *ErrorController) Error404() {
	this.SetTpl("error/404.tpl")
}

func (this *ErrorController) ErrorAjax401() {
	this.TplName = "error/ajax_401.tpl"
}

func (this *ErrorController) ErrorAjax403() {
	this.TplName = "error/ajax_403.tpl"
}

func (this *ErrorController) ErrorAjax404() {
	this.TplName = "error/ajax_404.tpl"
}
