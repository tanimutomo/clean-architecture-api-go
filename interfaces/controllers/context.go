package controllers

type Context interface {
	Param(string) string
	Bind(interface{}) error
	BindHeader(interface{}) error
	JSON(int, interface{})
	Abort()
}
