package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddCountry(country *models.Countries) error {
	tx := db.Data.Create(country)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AddLanguage(language *models.Languages) error {
	tx := db.Data.Create(language)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AddSysMessage(message *models.SysMessage) error {
	tx := db.Data.Create(message)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AddCurrency(currency *models.Currency) error {
	tx := db.Data.Create(currency)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AddTest(test *models.Test) error {
	tx := db.Data.Table("test").Create(test)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AddAgent(agents *models.Agents) error {
	tx := db.Data.Create(agents)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AddAccount(account *models.AccountAgent) error {
	tx := db.Data.Create(account)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}
