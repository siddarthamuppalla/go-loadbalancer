package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {

	serversURLList := unmarshalData("config.yml")
	var serverList []Server

	for _, url := range serversURLList {
		serverList = append(serverList, NewServer(url))
	}

	lb := NewLoadBalancer("8002", serverList)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	http.HandleFunc("/", handleRedirect)

	fmt.Printf("serving requests at 'localhost: %s' \n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}

func unmarshalData(filename string) []string {

	type T struct {
		Servers []string `yaml:"servers"`
	}

	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	var config T // empty slice for unmarshalling the data

	if err := yaml.Unmarshal(f, &config); err != nil {
		log.Fatal(err)
	}

	serversURLList := config.Servers
	return serversURLList
}
