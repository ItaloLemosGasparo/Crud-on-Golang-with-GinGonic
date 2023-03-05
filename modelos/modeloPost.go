package modelos

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Nome string
	Cod  int
}
