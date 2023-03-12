package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddTransfer(transfer *models.Transfers) error {
	tx := db.Data.Table("transfers").Create(transfer)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetTransfer(pagination *models.Pagination) ([]*models.Transfers, error) {
	var Transfer []*models.Transfers
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("transfers").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Transfers{}).Find(&Transfer)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return Transfer, nil
}

func (r Repository) UpdateTransfers(transfer *models.Transfers) error {

	var transfers *models.Transfers
	query := db.Data.Table("transfers").Where("id=?", transfer.ID).Find(&transfers)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}

	tx := db.Data.Table("transfers").Model(&models.Transfers{}).Where("id = ?", transfer.ID).Updates(models.Transfers{Entity: transfer.Entity, EntityId: transfer.EntityId, LangId: transfer.LangId, Value: transfer.Value})
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) DeleteTransfers(transfers *models.Transfers) error {
	log.Println(transfers.ID)

	query := db.Data.Table("transfers").Where("id =?", transfers.ID).Delete(transfers)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r Repository) TotalPageTransfer(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("transfers").Count(&length)
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
