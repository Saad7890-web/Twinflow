package proxy

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/Saad7890-web/twinflow/internal/storage"
	"github.com/Saad7890-web/twinflow/pkg/models"
)


type Proxy struct {
	target *url.URL
	proxy *httputil.ReverseProxy
	store  *storage.FileStore
}

func NewProxy(target string, store *storage.FileStore) (*Proxy, error){
	targetURL, err := url.Parse(target)

	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(targetURL)
	return &Proxy{
		target: targetURL,
		proxy:  rp,
		store:  store,
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


	record := models.TrafficRecord{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Method:    r.Method,
		Path:      r.URL.Path,
		Headers:   map[string]string{},
		RequestBody: string(requestBody),

		ResponseStatus: recorder.statusCode,
		ResponseBody:   recorder.body.String(),

		LatencyMs: duration.Milliseconds(),
	}

	for k, v := range r.Header {
		if len(v) > 0 {
		record.Headers[k] = v[0]
		}
	}

	p.store.Save(record)
	
}