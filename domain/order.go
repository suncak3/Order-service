package domain

type Order struct {
	OrdetID    uint `json:"order_id" db:"order_id"`
	ProductID  uint `json:"product_id" db:"product_id"`
	Quantities int  `json:"quantities" db:"quantities"`
}
