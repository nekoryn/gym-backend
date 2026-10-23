package domain

type Equipment struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Category    string `json:"category"`
	ImageURL    string `json:"image_url"`
	Description string `json:"description"`
}

