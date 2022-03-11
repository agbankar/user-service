package model

type User struct {
	Name string
	Id   int `json:"Id" validate:"code=CODE101" `
}
type UserBonus struct {
	UserID string  `json:"userID"`
	Bonus  float32 `json:"bonus"`
}
