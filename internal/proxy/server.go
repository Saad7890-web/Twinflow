package proxy

import (
	"log"
	"net/http"

	"github.com/Saad7890-web/twinflow/internal/storage"
)

func Start(listenAddr string, target string) error {

	store, err := storage.NewFileStore("captures/traffic.log")
	if err != nil {
		return err
	}

	proxy, err := NewProxy(target, store)
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