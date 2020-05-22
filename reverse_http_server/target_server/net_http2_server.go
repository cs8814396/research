package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

//https://www.mailgun.com/blog/http-2-cleartext-h2c-client-example-go/
func main() {
	// 在 8000 端口启动服务器
	// 确切地说，如何运行HTTP/1.1服务器。

	h2s := &http2.Server{}

	handler := http.HandlerFunc(sayhelloGolang)

	server := &http.Server{
		Addr:    "0.0.0.0:8000",
		Handler: h2c.NewHandler(handler, h2s),
	}

	server.ListenAndServe()

	/*
		srv := &http.Server{Addr: ":8000", Handler: h2c.NewHandler(sayhelloGolang, h2s)}
		// 用TLS启动服务器，因为我们运行的是http/2，它必须是与TLS一起运行。
		// 确切地说，如何使用TLS连接运行HTTP/1.1服务器。
		http2.ConfigureServer(srv, &http2.Server{})
		srv.ListenAndServeTLS("server.crt", "server.key")*/
}

func sayhelloGolang(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //解析参数，默认是不会解析的
	//fmt.Println("path", r.URL.Path)
	w.Write([]byte("Hello Golang"))
}
