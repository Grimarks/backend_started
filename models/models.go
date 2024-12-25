package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Product struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	SellerID    uint      `json:"seller_id"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	TransactionID uint      `json:"transaction_id"`
	PaymentStatus string    `json:"payment_status"`
	PaymentDate   time.Time `json:"payment_date"`
}

type Transaction struct {
	ID         uint      `gorm:"primaryKey"`
	ProductID  uint      `json:"product_id"`
	BuyerID    uint      `json:"buyer_id"`
	Quantity   int       `json:"quantity"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	Product    Product   `gorm:"foreignKey:ProductID"`
	Buyer      User      `gorm:"foreignKey:BuyerID"`
}
