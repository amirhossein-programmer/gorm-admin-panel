package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/models"
)

type Type struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CreateResponseType(typeMosel models.Type) Type {
	return Type{
		ID:   typeMosel.ID,
		Name: typeMosel.Name}
}
func FindType(id int, _type *models.Type) error {
	database.Database.Db.Find(&_type, "id = ?", id)
	if _type.ID == 0 {
		return errors.New("type is not found")
	}
	return nil
}
func CreateType(c *fiber.Ctx) error {
	var typeCreate models.Type

	if err := c.BodyParser(&typeCreate); err != nil {
		c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(typeCreate)
	responseType := CreateResponseType(typeCreate)
	return c.Status(200).JSON(responseType)
}
func GetTypes(c *fiber.Ctx) error {
	types := []models.Type{}
	database.Database.Db.Find(&types)
	responseTypes := []Type{}
	for _, _type := range types {
		responseType := CreateResponseType(_type)
		responseTypes = append(responseTypes, responseType)
	}
	return c.Status(200).JSON(responseTypes)
}
func GetType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var _type models.Type
	if err != nil {
		c.Status(400).JSON(err.Error())
	}
	if err := FindType(id, &_type); err != nil {
		c.Status(400).JSON(err.Error())
	}
	reasponseType := CreateResponseType(_type)
	return c.Status(200).JSON(reasponseType)
}
func UpdateType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var _type models.Type
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := FindType(id, &_type); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateType struct {
		Name string `json:"name"`
	}
	var updatedata UpdateType
	if err := c.BodyParser(&updatedata); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	_type.Name = updatedata.Name
	database.Database.Db.Save(&_type)
	responseTypes := CreateResponseType(_type)
	return c.Status(200).JSON(responseTypes)
}
func DeleteType(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var _type models.Type
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := FindType(id, &_type); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&_type).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}
	return c.Status(200).SendString("successfully deleted Type")
}

// done
