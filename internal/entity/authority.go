package entity

type Authority struct {
	ID   string `gorm:"primaryKey"`
	Name string
}
