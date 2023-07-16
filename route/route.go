package route

import (
	"net/http"

	"github.com/etnk125/go-webserver-merchant-management/internal/merchant"
	"github.com/labstack/echo/v4"
)

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

func Init(e *echo.Echo) {
	repo := merchant.NewMerchantRepository()
	service := merchant.NewMerchantService(repo)
	controller := merchant.NewController(service)

	routes := []Route{
		{Method: http.MethodGet, Path: "/", Handler: controller.HealthCheck},
		{Method: http.MethodGet, Path: "/health", Handler: controller.HealthCheck},
		{Method: http.MethodPost, Path: "/merchant", Handler: controller.RegisterMerchant},
		// {Method: http.MethodGet, Path: "/merchant/:merchant_id", Handler: controller.GetMerchantInfo, Middlewares: []echo.MiddlewareFunc{controller.AuthMiddleware}},
		// {Method: http.MethodPut, Path: "/merchant/:merchant_id", Handler: controller.UpdateMerchantInfo},
		// {Method: http.MethodPost, Path: "/merchant/:merchant_id/product", Handler: controller.AddProduct},
		// {Method: http.MethodGet, Path: "/merchant/:merchant_id/product/all", Handler: controller.GetProducts},
		// {Method: http.MethodPost, Path: "/product/:product_id", Handler: controller.BuyProduct},
		// {Method: http.MethodGet, Path: "/sell-report/:date", Handler: controller.GetSellReport},
	}

	for _, route := range routes {
		e.Add(route.Method, route.Path, route.Handler, route.Middlewares...)
	}
}

// | Method | Route                                | Body [JSON]        | Auth  | Description                    |
// | ------ | ------------------------------------ | ------------------ | ----- | ------------------------------ |
// | GET    | `/`                                  |                    | False | Health check                   |
// | GET    | `/health`                            |                    | False | Health check                   |
// | POST   | `/merchant`                          | name, bank_account | False | Register merchant              |
// | GET    | `/merchant/:merchant_id`             | name, bank_account | True  | Get merchant info              |
// | PUT    | `/merchant/:merchant_id`             | name, bank_account | True  | Update merchant info           |
// | POST   | `/merchant/:merchant_id/product`     | name, price        | True  | Add product to that merchant   |
// | GET    | `/merchant/:merchant_id/product/all` |                    | True  | Get products of that merchant  |
// | POST   | `/product/:product_id`               | quantity           | False | Buy product from that merchant |
// | GET    | `/sell-report/:date`                 |                    | True  | Get sell report                |
