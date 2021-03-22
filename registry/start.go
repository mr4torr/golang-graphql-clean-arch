package registry

import "gorm.io/gorm"

func Start(db *gorm.DB) Registry {
	return &registry{db}
}