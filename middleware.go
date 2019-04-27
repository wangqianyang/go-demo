package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)
/**
简单日志中间件
 */
func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		fmt.Fprint(w,"logging开始时间",start)
		next.ServeHTTP(w, r)
		log.Printf("Comleted %s in %v", r.URL.Path, time.Since(start))
		fmt.Fprint(w,"logging程序执行时间",time.Since(start))
	})
}
/**
简单认证中间件
 */
func authHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintln(w, "通过oauth认证...")
		next.ServeHTTP(w, r)

	})
}