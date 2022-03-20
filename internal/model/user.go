package model

type User struct {
	Name string `json:"Name" validate:"code=CODE101"`
	Id   int    `json:"Id" validate:"code=CODE102"`
}
type UserBonus struct {
	UserID string  `json:"userID"`
	Bonus  float32 `json:"bonus"`
}
