package main

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	port            string
	servers         []Server
	roundRobinCount int
}

func NewLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		servers:         servers,
		roundRobinCount: 0,
	}
}

func (lb *LoadBalancer) getNextAvailableServer() Server {

	server := lb.servers[lb.roundRobinCount%len(lb.servers)]

	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}

	lb.roundRobinCount++
	return server

}

func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {

	// get the next available server (target server)
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to: %q\n", targetServer.GetURL())

	// redirect to the servers url
	targetServer.Serve(rw, req)
}
