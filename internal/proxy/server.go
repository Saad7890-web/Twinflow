package proxy

import (
	"log"
	"net/http"
)

func Start(listenAddr string, target string) error {

	proxy, err := NewProxy(target)
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:    listenAddr,
		Handler: proxy,
	}

	log.Println("Proxy listening on", listenAddr)
	log.Println("Forwarding to", target)

	return server.ListenAndServe()
}