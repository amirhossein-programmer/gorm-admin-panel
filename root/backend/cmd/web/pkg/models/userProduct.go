package models

type UserProduct struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	UserRefer    int     `json:"userRefer"`
	User         User    `gorm:"foreignKey:UserRefer"`
	ProductRefer int     `json:"productRefer"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
}
