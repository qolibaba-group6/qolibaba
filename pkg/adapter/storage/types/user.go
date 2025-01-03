package types

type User struct {
	Model
	FirstName string
	LastName  string
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	IsAdmin   bool   `gorm:"default:false"`
	Status    uint8
	Role      string 
}
