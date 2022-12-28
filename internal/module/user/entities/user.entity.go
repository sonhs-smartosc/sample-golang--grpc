package entities

type User struct {
	Id    uint64 `json:"id" gorm:"primaryKey,autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Age   uint64 `json:"age"`
}
