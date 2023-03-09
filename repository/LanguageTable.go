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

func (r *Repository) GetLanguages(offset int, limit int) ([]*models.Languages, error) {
	var language []*models.Languages
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("languages").Limit(limit).Offset(offset).Find(&language)

	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return language, nil
}

func (r Repository) UpdateLanguage(language *models.Languages) error {

	var languages *models.Languages
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

	tx := db.Data.Model(&models.Languages{}).Where("id = ?", language.ID).Updates(models.Languages{Name: language.Name, Icon: language.Icon})
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
	tx := db.Data.Where("id", language.ID).Table("languages").Scan(&language)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if language.Active == true {
		language.Active = false
	} else {
		language.Active = true
	}
	tx = db.Data.Where("id", language.ID).Table("languages").Update("active", language.Active)
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
