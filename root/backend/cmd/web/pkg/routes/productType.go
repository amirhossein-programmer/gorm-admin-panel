package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/models"
)

type ProductType struct {
	ID      uint    `json:"id"`
	Product Product `json:"product"`
	Type    Type    `json:"type"`
}

func CreateResponseProductType(productTypeModel models.ProductType, _type Type, product Product) ProductType {
	return ProductType{ID: productTypeModel.ID, Product: product, Type: _type}
}

// CreateResponseProductType done next to createProductType method

func CreateProductType(c *fiber.Ctx) error {
	var productType models.ProductType
	if err := c.BodyParser(&productType); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	var product models.Product
	if err := FindProduct(productType.ProductRefer, &product); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	var _type models.Type
	if err := FindType(productType.TypeRefer, &_type); err != nil {
		return c.Status(404).JSON(err.Error())
	}
	database.Database.Db.Create(&productType)
	responseProduct := CreateResponseProduct(product)
	responseType := CreateResponseType(_type)
	responseProductType := CreateResponseProductType(productType, responseType, responseProduct)
	return c.Status(200).JSON(responseProductType)
}

func GetProductTypes(c *fiber.Ctx) error {
	poductTypes := []models.ProductType{}
	database.Database.Db.Find(&poductTypes)
	responseProductTypes := []ProductType{}
	for _, productType := range poductTypes {
		var product models.Product
		var _type models.Type
		database.Database.Db.Find(&product, "id = ?", productType.ProductRefer)
		database.Database.Db.Find(&_type, "id = ?", productType.TypeRefer)
		database.Database.Db.Find(&_type)
		responseProductType := CreateResponseProductType(productType, CreateResponseType(_type), CreateResponseProduct(product))
		responseProductTypes = append(responseProductTypes, responseProductType)
	}
	return c.Status(200).JSON(responseProductTypes)
}
func FindProductType(id int, ProductTypeModel *models.ProductType) error {
	database.Database.Db.Find(&ProductTypeModel, "id = ?", id)
	if ProductTypeModel.ID == 0 {
		return errors.New("product type is not found")
	}
	return nil
}
func GetProductType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(500).JSON(err.Error())
	}
	var productType models.ProductType
	if err := FindProductType(id, &productType); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	var product models.Product
	var _type models.Type

	database.Database.Db.First(&product, productType.ProductRefer)
	database.Database.Db.Find(&_type, "id = ?", productType.TypeRefer)

	responseProduct := CreateResponseProduct(product)
	responseType := CreateResponseType(_type)
	responseProductType := CreateResponseProductType(productType, responseType, responseProduct)

	return c.Status(200).JSON(responseProductType)
}
