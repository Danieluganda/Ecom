package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Create instances of all microservices
	productService := NewProductService()
	orderService := NewOrderService()
	userService := NewUserService()
	paymentService := NewPaymentService()
	cartService := NewCartService()
	recommendationService := NewRecommendationService()
	shippingService := NewShippingService()
	notificationService := NewNotificationService()

	// Create a goroutine for each microservice to handle incoming requests
	go productService.Start()
	go orderService.Start()
	go userService.Start()
	go paymentService.Start()
	go cartService.Start()
	go recommendationService.Start()
	go shippingService.Start()
	go notificationService.Start()

	// Set up HTTP endpoints to route requests to the appropriate microservice
	http.HandleFunc("/products", productService.HandleRequest)
	http.HandleFunc("/orders", orderService.HandleRequest)
	http.HandleFunc("/users", userService.HandleRequest)
	http.HandleFunc("/payments", paymentService.HandleRequest)
	http.HandleFunc("/cart", cartService.HandleRequest)
	http.HandleFunc("/recommendations", recommendationService.HandleRequest)
	http.HandleFunc("/shipping", shippingService.HandleRequest)
	http.HandleFunc("/notifications", notificationService.HandleRequest)

	// Start the HTTP server in a separate goroutine
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal(err)
		}
	}()

	fmt.Println("Microservices are running...")

	// Gracefully handle termination signals to shut down all microservices
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down all microservices...")

	// Stop all microservices gracefully
	productService.Stop()
	orderService.Stop()
	userService.Stop()
	paymentService.Stop()
	cartService.Stop()
	recommendationService.Stop()
	shippingService.Stop()
	notificationService.Stop()

	fmt.Println("Microservices have been stopped.")
}
