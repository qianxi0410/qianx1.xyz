package dao

type User struct {
	ID       string `json:"id" gorm:"primaryKey" redis:"id" form:"id"`
	Name     string `json:"name" gorm:"name" redis:"name" form:"name"`
	Password string `json:"password" gorm:"password" redis:"password" form:"password"`
}
