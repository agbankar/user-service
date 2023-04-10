package model

type User struct {
	Id        int    `json:"id" validate:"code=CODE102"`
	FirstName string `json:"firstName" validate:"code=CODE101"`
	LastName  string `json:"lastName"`
	Address   string `json:"address"`
	City      string `json:"city"`
}
type UserBonus struct {
	UserID string  `json:"userID"`
	Bonus  float32 `json:"bonus"`
}

type Persons struct {
	PersonId  string `gorm:"primaryKey"`
	LastName  string
	FirstName string
	Address   string
	City      string
}
