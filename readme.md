##demo代码结构
main.go             --> 应用入口，web服务启动，请求路由，中间件设置

controller.go       --> 请求处理逻辑具体实现代码

middleware.go       --> 中间件实现代码，简单模拟实现了日志，鉴权功能

controller_test.go  --> 单元测试，基于net/http/httptest，testing实现

goods.go            --> DB层代码，简单实现mongodb的CRUD功能
