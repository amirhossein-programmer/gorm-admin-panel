package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/database"
	"github.com/amirhossein-programmer/fiber-api/cmd/web/pkg/models"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
	Gender    string `json:"gender"`
}

func FindUser(id int, user *models.User) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("user not found")
	}
	return nil
}
func CreateResponseUser(UserModel models.User) User {
	return User{
		ID:        UserModel.ID,
		UserName:  UserModel.UserName,
		Phone:     UserModel.Phone,
		FirstName: UserModel.FirstName,
		LastName:  UserModel.LastName,
		Email:     UserModel.Email,
		Age:       UserModel.Age,
		Gender:    UserModel.Gender,
	}
}
func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("please ensure that the id is integer")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	responsUser := CreateResponseUser(user)
	return c.Status(200).JSON(responsUser)
}
func LoginUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}
func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("please ensure that the id is integer")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	type UpdateUser struct {
		UserName  string `json:"user_name"`
		Phone     string `json:"phone"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Age       uint   `json:"age"`
		Gender    string `json:"gender"`
	}
	var updateData UpdateUser
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}
	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName
	user.UserName = updateData.UserName
	user.Email = updateData.Email
	user.Phone = updateData.Phone
	user.Age = updateData.Age
	user.Gender = updateData.Gender
	database.Database.Db.Save(&user)
	responseUsers := CreateResponseUser(user)
	return c.Status(200).JSON(responseUsers)
}
func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return c.Status(400).JSON("please ensure that the id is integer")
	}
	if err := FindUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	if err := database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).SendString("successfully deleted user")
}

// done
