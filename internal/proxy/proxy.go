package proxy

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)


type Proxy struct {
	target *url.URL
	proxy *httputil.ReverseProxy
}

func NewProxy(target string) (*Proxy, error){
	targetURL, err := url.Parse(target)

	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(targetURL)
	return &Proxy{
		target: targetURL,
		proxy:  rp,
	}, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request){
	start := time.Now()

	var requestBody []byte
	if r.Body != nil {
		requestBody, _ = io.ReadAll(r.Body)
		r.Body = io.NopCloser(bytes.NewBuffer(requestBody))
	}

	log.Println("Incoming request:", r.Method, r.URL.Path)

	recorder := &responseRecorder{
		ResponseWriter: w,
		body:           &bytes.Buffer{},
		statusCode:     http.StatusOK,
	}

	p.proxy.ServeHTTP(recorder, r)

	duration := time.Since(start)

	log.Println("Response status:", recorder.statusCode)
	log.Println("Latency:", duration)

	log.Println("Request Body:", string(requestBody))
	log.Println("Response Body:", recorder.body.String())


	
}