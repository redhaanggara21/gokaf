package config

import "gorm.io/gorm"

func connectDatabase() *gorm.DB {
	db, err := gorm.Open(
		postgresql.Open("postgre:")
	)
}
