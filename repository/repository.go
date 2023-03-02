package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	Connection *gorm.DB
}

func NewRepository(conn *gorm.DB) *Repository {
	return &Repository{Connection: conn}
}

// не нужный функционал
//func (r Repository) Migrate() error {
//	err := db.Data.AutoMigrate(
//		&models.Countries{},
//		&models.Languages{},
//		&models.Currency{},
//		&models.SysMessage{},
//		&models.Test{},
//		&models.Agents{},
//		&models.AccountAgent{},
//	)
//
//	if err != nil {
//		log.Println(err)
//		return err
//	}
//	return nil
//}
