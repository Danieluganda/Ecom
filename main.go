
package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/gorilla/mux"

	
	"github.com/Danieluganda/Ecom/user-service"
	"github.com/Danieluganda/Ecom/authentication-service"
	"github.com/Danieluganda/Ecom/roduct-service"
	"github.com/Danieluganda/Ecom/nventory-service"
	"github.com/Danieluganda/Ecom//cart-service"
	"github.com/Danieluganda/Ecom/order-service"
	"github.com/Danieluganda/Ecom/payment-service"
	"github.com/Danieluganda/Ecom/notification-service"
	"github.com/Danieluganda/Ecom/api-gateway"
	"github.com/Danieluganda/Ecom/service-discovery"
	"github.com/Danieluganda/Ecom/logging"
	"github.com/Danieluganda/Ecom/distributed-tracing"
	
)

func main() {
	router := mux.NewRouter()

	// Initialize Services
	userSvc := user.NewUserService()
	authSvc := authentication.NewAuthenticationService()
	productSvc := product.NewProductService()
	inventorySvc := inventory.NewInventoryService()
	cartSvc := cart.NewCartService()
	orderSvc := order.NewOrderService()
	paymentSvc := payment.NewPaymentService()
	notificationSvc := notification.NewNotificationService()

	// Initialize Middleware
	logger := logging.NewLoggerMiddleware()
	tracer := tracing.NewTracingMiddleware()

	// Initialize API Gateway
	apiGateway := gateway.NewAPIGateway(router, logger, tracer)

	// Register Service Endpoints
	apiGateway.RegisterEndpoint("/users", http.MethodGet, userSvc.GetUsers)
	apiGateway.RegisterEndpoint("/authenticate", http.MethodPost, authSvc.Authenticate)
	apiGateway.RegisterEndpoint("/products", http.MethodGet, productSvc.GetProducts)
	apiGateway.RegisterEndpoint("/inventory", http.MethodGet, inventorySvc.GetInventory)
	apiGateway.RegisterEndpoint("/carts", http.MethodGet, cartSvc.GetCarts)
	apiGateway.RegisterEndpoint("/orders", http.MethodGet, orderSvc.GetOrders)
	apiGateway.RegisterEndpoint("/payments", http.MethodGet, paymentSvc.GetPayments)
	apiGateway.RegisterEndpoint("/notifications", http.MethodGet, notificationSvc.GetNotifications)

	// Initialize Service Discovery
	discovery := serviceDiscovery.NewServiceDiscovery()
