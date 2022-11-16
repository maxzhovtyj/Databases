package orm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type StorageConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewORMClient(cfg *StorageConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.Username,
		cfg.Password,
		cfg.Database,
		cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
