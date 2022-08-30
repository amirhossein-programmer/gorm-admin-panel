package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/models"
)

type UserProduct struct {
	ID      uint    `json:"id"`
	User    User    `json:"user"`
	Product Product `json:"product"`
}

func createResponseUserProduct(userProductModel models.UserProduct, product Product, user User) UserProduct {
	return UserProduct{ID: userProductModel.ID, Product: product, User: user}
}
func CreateUserProduct(c *fiber.Ctx) error {
	var userProduct models.UserProduct
	if err := c.BodyParser(&userProduct); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var user models.User
	if err := FindUser(userProduct.UserRefer, &user); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	var product models.Product
	if err := FindProduct(userProduct.ProductRefer, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	database.Database.Db.Create(&userProduct)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	respnseUserProduct := createResponseUserProduct(userProduct, responseProduct, responseUser)
	return c.Status(200).JSON(respnseUserProduct)
}
