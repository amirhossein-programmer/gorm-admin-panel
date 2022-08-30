package models

type Order struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	UserRefer    int     `json:"User_id"`
	User         User    `gorm:"foreignKey:UserRefer"`
	ProductRefer int     `json:"Product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
}
