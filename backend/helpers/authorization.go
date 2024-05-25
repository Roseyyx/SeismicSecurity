package helpers

import (
	"main/backend/models"
)

func Login(password, file string) bool {
	bytes := OpenFile(file, password)
	if bytes == nil {
		return false
	}
	return true
}

func Register(password string) {
	file := new(models.File)

	file.Filename = "database"

	CreateFile(file.Filename, password)
}
