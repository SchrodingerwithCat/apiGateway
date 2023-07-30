// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	handler "hertz.demo/biz/handler"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
}
