// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	login "admin/biz/router/login"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	login.Register(r)

}
