# GoMVC

go get github.com/greatbsky/GoMVC
go get github.com/gorilla/websocket
go get github.com/astaxie/beego

直接运行：
E:\MyWork7\GoMVC>go run main.go
2018/02/06 16:02:01.605 [I] [asm_386.s:1635] http server Running on http://:8080/
2018/02/06 16:02:11.746 [D] [server.go:2619] 127.0.0.1 - - [06/Feb/2018 04:02:11] "GET / HTTP/1.1 200 0" 0.007000  Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36
/static/js/reload.min.js
2018/02/06 16:02:11.891 [D] [server.go:2619] 127.0.0.1 - - [06/Feb/2018 04:02:11] "GET /static/js/reload.min.js HTTP/1.1 404 0" 0.001000 http://127.0.0.1:8080/ Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36

编译：
E:\MyWork7\GoMVC>go build main.go
E:\MyWork7\GoMVC>main.exe
2018/02/06 16:03:44.829 [I] [asm_386.s:1635] http server Running on http://:8080
