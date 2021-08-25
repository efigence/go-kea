package go_kea

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	DSN string
}

type Kea struct {
	db *gorm.DB
}

func New(cfg Config) (*Kea, error) {
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	k := Kea{db: db}
	return &k, nil
}
