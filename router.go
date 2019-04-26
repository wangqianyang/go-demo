package main

import (
	"fmt"
	"net/http"
	"reflect"
)


func (p *ControllerRegistor) Add(pattern string, c ControllerInterface) {

	//添加路由
	t := reflect.TypeOf(c).Elem()
	route := &controllerInfo{}
	route.url = pattern
	route.controllerType = t
	p.routers = append(p.routers, route)

}

// 路由
func (p *ControllerRegistor) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var find bool
	requestPath := r.URL.Path

	// fmt.Println(requestPath)

	// 查找路由注册表
	for _, route := range p.routers {

		if requestPath == route.url {
			vc := reflect.New(route.controllerType)
			method := vc.MethodByName("Do")
			method.Call(nil)
			find = true
			fmt.Fprintf(w, "Hello "+route.controllerType.Name())
			break
		}
	}

	//没有找到，返回404
	if find == false {
		http.NotFound(w, r)
	}
}
