package main

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"

	"github.com/go-resty/resty/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/publicsuffix"
)

func main() {
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}

	var t http.RoundTripper

	t1 := http.DefaultTransport.(*http.Transport).Clone()
	t1.DialContext = dialer.DialContext
	t1.MaxIdleConns = 200
	//t1.MaxConnsPerHost = 500
	t1.MaxIdleConnsPerHost = 100
	t1.ForceAttemptHTTP2 = true
	t1.TLSHandshakeTimeout = 10 * time.Second
	t1.ExpectContinueTimeout = 1 * time.Second
	t1.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// HTTP/1 -> HTTP/2
	// t2, err := http2.ConfigureTransports(t1)
	_ = http2.ConfigureTransport(t1)
	t = t1

	// if err == nil {
	// 	fmt.Println("h2")
	// 	t2.AllowHTTP = true
	// 	t2.DialTLS = func(netw, addr string, cfg *tls.Config) (net.Conn, error) {
	// 		return net.Dial(netw, addr)
	// 	}
	// 	t = t2
	// } else {
	// 	fmt.Println("t1")
	// 	t = t1
	// }

	cookieJar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	httpClient := &http.Client{
		Timeout:   10 * time.Second,
		Transport: t,
		Jar:       cookieJar,
	}

	client := resty.NewWithClient(httpClient)

	req := client.R()

	req.SetHeader("X-Request-By", "local_h2c_client")

	resp, err := req.Get("https://http2.golang.org/reqinfo")

	if err != nil {
		if errors.Is(err, http2.ErrNoCachedConn) {
			fmt.Println("h2 error", err)
		}

		panic(err)
	}

	fmt.Println(resp.RawResponse.Proto)
	fmt.Println(resp.RawResponse.StatusCode)
}
