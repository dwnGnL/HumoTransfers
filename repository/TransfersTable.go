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

func (r *Repository) GetTransfer(offset int, limit int) ([]*models.Transfers, error) {
	var transfers []*models.Transfers
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("transfers").Limit(limit).Offset(offset).Find(&transfers)

	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return transfers, nil
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
