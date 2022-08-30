package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/routes"
)

func SetupRoutes(app *fiber.App) {

	home := app.Group("/home")
	user := home.Group("/user")
	product := home.Group("/product")
	productType := product.Group("/type")
	order := home.Group("/order")
	_type := home.Group("/type")
	userProduct := user.Group("/product")

	// setup routes users
	user.Post("/create", routes.CreateUser)       // home/user/create
	user.Put("/update/:id", routes.UpdateUser)    // home/user/update/:id
	user.Delete("/delete/:id", routes.DeleteUser) // home/user/delete/:id
	user.Get("/", routes.GetUsers)                // home/user
	user.Get("/:id", routes.GetUser)              // home/user/:id

	// setup routes product
	product.Post("/create", routes.CreateProduct)       // home/product/create
	product.Put("/update/:id", routes.UpdateProduct)    // home/product/update/:id
	product.Delete("/delete/:id", routes.DeleteProduct) // home/product/delete/:id
	product.Get("/", routes.GetProducts)                // home/product
	product.Get("/:id", routes.GetProduct)              // home/product/:id
	// setup routes type
	_type.Post("/create", routes.CreateType)       // home/) // home/type/create
	_type.Put("/update/:id", routes.UpdateType)    // home/type/update/:id
	_type.Delete("/delete/:id", routes.DeleteType) // home/type/delete/:id
	_type.Get("/", routes.GetTypes)                // home/type/
	_type.Get("/:id", routes.GetType)              // home/type/:id
	// setup routes productType
	productType.Post("/create", routes.CreateProductType) // home/product/type/create
	productType.Get("/", routes.GetProductTypes)          // home/product/type/
	productType.Get("/:id", routes.GetProductType)        // home/poduct/type/:id
	// setup routes user product
	userProduct.Post("/create", routes.CreateUserProduct) // home/user/poduct/create
	// setup routes order
	order.Post("/create", routes.CreateOrder) // home/order/create
	order.Get("/", routes.GetOrders)          // home/order
	order.Get("/:id", routes.GetOrder)        // home/order/:id
}
