package models

type ProductType struct {
	ID           uint    `json:"id"`
	ProductRefer int     `json:"product_id"`
	Product      Product `gorm:"foreignKey:ProductRefer"`
	TypeRefer    int     `json:"type_id"`
	Type         Type    `gorm:"foreignKey:TypeRefer"`
}
