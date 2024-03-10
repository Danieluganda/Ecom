package main

import (
	"errors"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type ServiceRegistry struct {
	mutex     sync.RWMutex
	services  map[string]string
	endpoints map[string][]string
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{
		services:  make(map[string]string),
		endpoints: make(map[string][]string),
	}
}

func (sr *ServiceRegistry) RegisterService(serviceName, serviceAddress string, endpoints []string) {
	sr.mutex.Lock()
	defer sr.mutex.Unlock()

	sr.services[serviceName] = serviceAddress
	sr.endpoints[serviceName] = endpoints
}

func (sr *ServiceRegistry) ResolveService(serviceName string) (string, []string, error) {
	sr.mutex.RLock()
	defer sr.mutex.RUnlock()

	serviceAddress, ok := sr.services[serviceName]
	if !ok {
		return "", nil, errors.New("service not found")
	}

	endpoints, ok := sr.endpoints[serviceName]
	if !ok {
		return "", nil, errors.New("endpoints not found for service")
	}

	return serviceAddress, endpoints, nil
}

func main() {
	serviceRegistry := NewServiceRegistry()

	serviceName := "my-service"
	serviceAddress := "127.0.0.1:8080"
	endpoints := []string{"/api/endpoint1", "/api/endpoint2"}

	serviceRegistry.RegisterService(serviceName, serviceAddress, endpoints)

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		service, _, err := serviceRegistry.ResolveService(serviceName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// Forward the request to the service
		reverseProxy := NewReverseProxy(service)
		reverseProxy.ServeHTTP(w, req)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func NewReverseProxy(target string) *httputil.ReverseProxy {
	remote, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	return httputil.NewSingleHostReverseProxy(remote)
}
