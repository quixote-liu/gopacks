package model

type Student struct {
	ID          string
	Name        string
	StudentCard StudentCard
}

type StudentCard struct {
	Number    string `gorm:"primaryKey"`
	StudentID string
}

