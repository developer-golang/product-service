package models

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

type ProductDetails struct {
	ID          int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string  `json:"title" gorm:"type:varchar(100)"`
	Price       float64 `json:"price" gorm:"type:decimal(10,2)"`
	Description string  `json:"description" gorm:"type:varchar(5000)"`
	Category    string  `json:"category" gorm:"type:varchar(200)"`
	Image       string  `json:"image" gorm:"type:varchar(300)"`
	Rating      Rating  `json:"rating" gorm:"serializer:json;type:jsonb"`
}
