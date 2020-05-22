package main

import (
	"bytes"
	"fmt"
	"golang.org/x/net/http2"
	"io/ioutil"
	"net/http"

	"crypto/tls"
	"net/http/httputil"
	//"time"
	"net"
)

var httpClient *http.Client

func init() {
	/*
		caCert, err := ioutil.ReadFile("server.crt")
		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM(caCert)
		tlsConfig := &tls.Config{RootCAs: caCertPool}

		client.Transport = &http2.Transport{TLSClientConfig: tlsConfig}*/
	/*tr := &http2.Transport{
		AllowHTTP: true, //充许非加密的链接
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	httpClient = &http.Client{Transport: tr}*/
	httpClient = &http.Client{
		Transport: &http2.Transport{
			// So http2.Transport doesn't complain the URL scheme isn't 'https'
			AllowHTTP: true,
			// Pretend we are dialing a TLS endpoint.
			// Note, we ignore the passed tls.Config
			DialTLS: func(network, addr string, cfg *tls.Config) (net.Conn, error) {
				return net.Dial(network, addr)
			},
		},
	}

}
func director(req *http.Request) {

	if req.URL.Scheme == "" {
		req.URL.Scheme = "http"
	}

	req.URL.Path = "/_internal/notAllowed"
	req.URL.Host = "127.0.0.1:14001"
	// check

	url := "http://127.0.0.1:8000"

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
