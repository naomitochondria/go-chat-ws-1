package domain

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Passord  string
}
