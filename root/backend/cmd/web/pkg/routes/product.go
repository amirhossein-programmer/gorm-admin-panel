package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/models"
)

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func CreateResponseProduct(productModel models.Product) Product {
	return Product{ID: productModel.ID, Name: productModel.Name, Price: productModel.Price}
}
func FindProduct(id int, product *models.Product) error {
	database.Database.Db.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("Product not found")
	}
	return nil
}
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}
	database.Database.Db.Find(&products)
	responseProducts := []Product{}
	for _, product := range products {
		responseProduct := CreateResponseProduct(product)
		responseProducts = append(responseProducts, responseProduct)
	}
	return c.Status(200).JSON(responseProducts)
}
func GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		return c.Status(400).JSON("please ensure the product ID is valid")
	}
	if err := FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
	}
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}
func UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		c.Status(400).JSON("please ensure that id is valid")
	}
	if err := FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
	}
	type UpdateProduct struct {
		Name  string  `json:"name"`
		Price float64 `json:"price"`
	}
	var updateData UpdateProduct
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	product.Name = updateData.Name
	product.Price = updateData.Price
	database.Database.Db.Save(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)
}
func DeleteProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var product models.Product
	if err != nil {
		c.Status(400).JSON("please ensure that id is valid")
	}
	if err := FindProduct(id, &product); err != nil {
		c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("successfully deleted product")
}

// done
