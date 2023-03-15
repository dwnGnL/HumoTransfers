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

func (r *Repository) GetCurrency(pagination *models.Pagination) ([]models.Currency, error) {
	var currency []models.Currency
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("currencies").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Currency{}).Find(&currency)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return currency, nil
}

func (r Repository) UpdateCurrency(currency *models.Currency) error {

	var currencies models.Languages
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

	tx := db.Data.Model(models.Currency{}).Where("id = ?", currency.ID).Updates(models.Currency{Name: currency.Name, Icon: currency.Icon})
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

func (r Repository) TotalPageCurrency(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("currencies").Count(&length)
	if query.Error != nil {
		log.Println(query.Error)
		return 0, query.Error
	}
	totalPage := length / limit
	if length%limit != 0 {
		totalPage++
	}
	return totalPage, nil
}
