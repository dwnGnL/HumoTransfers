package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddSysMessage(message *models.SysMessage) error {
	tx := db.Data.Create(message)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetSysMessage(offset int, limit int) ([]*models.SysMessage, error) {
	var message []*models.SysMessage
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("sys_messages").Limit(limit).Offset(offset).Find(&message)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return message, nil
}

func (r Repository) UpdateSysMessage(message *models.SysMessage) error {
	if message.Name == "" {
		query := db.Data.Table("sys_messages").Where("id =?", message.ID).Delete(message)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {

		tx := db.Data.Model(&models.SysMessage{}).Where("id = ?", message.ID).Updates(models.SysMessage{Name: message.Name})
		if tx.Error != nil {
			log.Println(tx.Error)
			return tx.Error
		}
	}
	return nil
}

func (r Repository) DeleteSysMessage(message *models.SysMessage) error {

	query := db.Data.Table("sys_messages").Where("id =?", message.ID).Delete(message)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r Repository) SysMessageStatus(message *models.SysMessage) error {
	tx := db.Data.Where("id", message.ID).Table("sys_messages").Scan(&message)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if message.Active == true {
		message.Active = false
	} else {
		message.Active = true
	}
	tx = db.Data.Where("id", message.ID).Table("sys_messages").Update("active", message.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}
