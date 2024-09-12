package db

import (
	"fmt"

	"github.com/sds-2/config"
	models "github.com/sds-2/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreSQL(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", cfg.PGDB.Host, cfg.PGDB.User, cfg.PGDB.Password, cfg.PGDB.Name, cfg.PGDB.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migration(db)
	return db
}

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.Item{})
}
