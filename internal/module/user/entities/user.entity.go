package entities

type User struct {
	Id       uint64 `json:"id" gorm:"primaryKey,autoIncrement"`
	Name     string `json:"name"`
	Age      uint64 `json:"age"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
