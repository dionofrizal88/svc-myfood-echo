package routes

import (
	"net/http"
	"svc-myfood-echo/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() *echo.Echo{
	e := echo.New()

	e.GET("/", func(c echo.Context)error{
		return c.String(http.StatusOK, "Hello this is echo!")
	})

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
      AllowOrigins: []string{"*"},
      AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
    }))  

	e.GET("/employee", controllers.FetchAllEmployee)
	e.POST("/employee", controllers.StoreEmployee)
	e.PUT("/employee/:id", controllers.UpdateEmployee)
	e.DELETE("/employee/:id", controllers.DeleteEmployee)

	e.GET("/products", controllers.FetchAllProduct)
	e.GET("/products/:uuid", controllers.FetchProductByUuid)
	e.POST("/products", controllers.StoreProduct)

	e.GET("/basket", controllers.FetchActiveBasket)
	e.POST("/basket/:uuid_product", controllers.StoreProductToBasket)
	e.DELETE("/basket/:uuid_product", controllers.DeleteProductToBasket)

	e.POST("/checkout", controllers.CheckoutBasketProduct)
	return e
}

// CORSRouterDecorator applies CORS headers to a mux.Router
// type CORSRouterDecorator struct {
// 	R *mux.Router
// }

// func (c *CORSRouterDecorator) ServeHTTP(rw http.ResponseWriter,
// 	req *http.Request) {
// 	if origin := req.Header.Get("Origin"); origin != "" {
// 		rw.Header().Set("Access-Control-Allow-Origin", origin)
// 		rw.Header().Set("Access-Control-Allow-Methods",
// 			"POST, GET, OPTIONS, PUT, DELETE")
// 		rw.Header().Set("Access-Control-Allow-Headers",
// 			"Accept, Accept-Language,"+
// 				" Content-Type, YourOwnHeader")
// 	}
// 	// Stop here if its Preflighted OPTIONS request
// 	if req.Method == "OPTIONS" {
// 		return
// 	}

// 	c.R.ServeHTTP(rw, req)
// }