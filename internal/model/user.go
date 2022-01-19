package model

type User struct {
	Name string
	Id   int
}
type UserBonus struct {
	UserID string  `json:"userID"`
	Bonus  float32 `json:"bonus"`
}
