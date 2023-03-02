package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddCurrency(currency *models.Currency) error {
	tx := db.Data.Create(currency)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetCurrency(offset int, limit int) ([]*models.Currency, error) {
	var currency []*models.Currency
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("currencies").Limit(limit).Offset(offset).Find(&currency)

	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return currency, nil
}

func (r Repository) UpdateCurrency(currency *models.Currency) error {

	var currencies *models.Languages
	query := db.Data.Where("id=?", currency.ID).Find(&currencies)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	if currency.Name == "" {
		currency.Name = currencies.Name
	}
	if currency.Icon == "" {
		currency.Icon = currencies.Icon
	}

	tx := db.Data.Model(&models.Currency{}).Where("id = ?", currency.ID).Updates(models.Currency{Name: currency.Name, Icon: currency.Icon})
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) DeleteCurrency(currency *models.Currency) error {
	query := db.Data.Table("currencies").Where("id =?", currency.ID).Delete(currency)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}
