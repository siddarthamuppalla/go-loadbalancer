package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	IsAlive() bool
	GetURL() *url.URL
	Serve(http.ResponseWriter, *http.Request)
}

type server struct {
	url          *url.URL
	reverseProxy *httputil.ReverseProxy
}

func (s *server) IsAlive() bool {
	return true
}

func (s *server) GetURL() *url.URL {
	return s.url
}

func (s *server) Serve(rw http.ResponseWriter, req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	s.reverseProxy.ServeHTTP(rw, req)
}

func NewServer(u string) *server {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		fmt.Printf("error: %v\n", err.Error())
		os.Exit(1)
	}

	return &server{
		url:          parsedUrl,
		reverseProxy: httputil.NewSingleHostReverseProxy(parsedUrl),
	}
}
