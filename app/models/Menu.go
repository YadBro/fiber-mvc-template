package models

type Menu struct {
	Id       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ParentId uint   `json:"parent_id" gorm:"null;index"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Urutan   uint   `json:"urutan" gorm:"not null;column:urutan"`
	Children []Menu `gorm:"foreignKey:ParentId;references:Id"`
}
