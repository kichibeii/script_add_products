package commons

import (
	"github.com/script_add_products/server/config"
	"gorm.io/gorm"
)

type Options struct {
	Config   *config.Configuration
	Database *gorm.DB
}
