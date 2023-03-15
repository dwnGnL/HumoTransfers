package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddLanguage(language *models.Languages) error {
	tx := db.Data.Create(language)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r *Repository) GetLanguages(pagination *models.Pagination) ([]models.Languages, error) {
	var Languages []models.Languages
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("languages").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Languages{}).Find(&Languages)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return Languages, nil
}

func (r Repository) UpdateLanguage(language *models.Languages) error {

	var languages models.Languages
	query := db.Data.Where("id=?", language.ID).Find(&languages)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	if language.Name == "" {
		language.Name = languages.Name
	}
	if language.Icon == "" {
		language.Icon = languages.Icon
	}

	tx := db.Data.Model(models.Languages{}).Where("id = ?", language.ID).Updates(models.Languages{Name: language.Name, Icon: language.Icon})
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r Repository) DeleteLanguage(language *models.Languages) error {

	query := db.Data.Table("languages").Where("id =?", language.ID).Delete(language)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r Repository) LanguageStatus(language *models.Languages) error {
	tx := db.Data.Where("id = ?", language.ID).Table("languages").Scan(language)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if language.Active == true {
		language.Active = false
	} else {
		language.Active = true
	}
	tx = db.Data.Where("id = ?", language.ID).Table("languages").Update("active", language.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) TotalPageLanguage(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}
	query := db.Data.Table("languages").Count(&length)
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
