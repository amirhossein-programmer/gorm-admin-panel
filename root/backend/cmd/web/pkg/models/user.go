package models

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserName  string `json:"user_name"`
	Phone     string `json:"phone"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Age       uint   `json:"age"`
	Gender    string `json:"gender"`
}
