package model

type User struct {
	Id        int      `json:"id" validate:"code=CODE102" gorm:"primaryKey"`
	FirstName string   `json:"firstName" validate:"code=CODE101"`
	LastName  string   `json:"lastName"`
	Address   string   `json:"address"`
	City      string   `json:"city"`
	Orders    []Orders `json:"orders" gorm:"foreignKey:Id"`
}
type UserBonus struct {
	UserID string  `json:"userID"`
	Bonus  float32 `json:"bonus"`
}
type Orders struct {
	Id      int `json:"id"`
	OrderId int ` json:"orderId" gorm:"primaryKey"`
}

func (u User) TableName() string {
	return "users"

}

func (u Orders) TableName() string {
	return "orders"

}
