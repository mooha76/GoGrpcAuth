package db

import (
	"fmt"

	"github.com/mooha76/GoGrpcAuth/config"
	model "github.com/mooha76/GoGrpcAuth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		model.RefreshSession{},
		model.OtpSession{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
