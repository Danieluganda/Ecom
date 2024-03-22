
package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/gorilla/mux"

	/*
	"github.com/your-package/user-service"
	"github.com/x/oauth2/authentication-service"
	"github.com/your-package/product-service"
	"github.com/your-package/inventory-service"
	"github.com/your-package/cart-service"
	"github.com/your-package/order-service"
	"github.com/your-package/payment-service"
	"github.com/your-package/notification-service"
	"github.com/your-package/api-gateway"
	"github.com/your-package/service-discovery"
	"github.com/your-package/logging"
	"github.com/your-package/distributed-tracing"
	*/
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
