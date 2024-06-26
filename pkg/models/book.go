package models

type Book struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Title  string `json:"title" validate:"required"`
	Author string `json:"author" validate:"required"`
	Desc   string `json:"desc" validate:"required"`
}