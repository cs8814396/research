package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

var httpClient *http.Client

func init() {

	httpClient = &http.Client{}

}
func director(req *http.Request) {

	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	req.URL.Path = "/_internal/notAllowed"
	req.URL.Host = "127.0.0.1:14001"
	// check

	url := "http://127.0.0.1:8080"

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte("hello")))
	if err != nil {
		panic("why httpReq fail" + err.Error())
	}

	httpResp, err := httpClient.Do(httpReq)
	if err != nil {
		panic("why do fail " + err.Error())
	}
	defer httpResp.Body.Close()
	//下面的代码注释打开了可以看返回的内容，但就不能json反序列化了
	_, err = ioutil.ReadAll(httpResp.Body)
	if err == nil && httpResp.StatusCode == 200 {

	}

}

func modifyResponse(resp *http.Response) error {

	return nil
}

func main() {

	accessGwProxy := &httputil.ReverseProxy{Director: director, ModifyResponse: modifyResponse}
	http.Handle("/", accessGwProxy)
	http.HandleFunc("/_internal/notAllowed", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `{"errcode":20503,"errmsg":"no privilege"}`)
	})
	http.ListenAndServe(":14001", nil)
	// http.ListenAndServe(":24001", nil)
}
