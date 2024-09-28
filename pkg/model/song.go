package model

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Group       string
	Song        string
	ReleaseDate string
	Text        string
	Link        string
}
