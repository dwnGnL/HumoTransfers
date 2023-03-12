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

//func (r *Repository) GetCountries(offset int, limit int) ([]*models.Countries, error) {
//	var countries []*models.Countries
//if limit == 0 {
//	limit = 10
//}
//
//query := db.Data.Table("countries")
////if filter!="" {
////	query = query.Where("name like ?","%"+filter+"%")
////}
//
//tx := query.Limit(limit).Offset(offset).Find(&countries)
//if tx.Error != nil {
//	log.Println(tx.Error)
//	return nil, tx.Error
//}
//return countries, nil

func (r *Repository) GetCountries(pagination *models.Pagination) ([]*models.Countries, error) {
	var countries []*models.Countries
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuider := db.Data.Table("countries").Limit(pagination.Limit).Offset(offset)
	result := queryBuider.Model(&models.Countries{}).Find(&countries)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return countries, nil
}

func (r Repository) UpdateCountries(country *models.Countries) error {

	var countries *models.Countries
	query2 := db.Data.Where("id=?", country.ID).Find(&countries)
	if query2.Error != nil {
		log.Println(query2.Error)
		return query2.Error
	}

	log.Println(1, country, 2, countries)
	tx := db.Data.Model(&models.Countries{}).Where("id = ?", country.ID).Updates(models.Countries{Name: country.Name, Icon: country.Icon, Active: country.Active})
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r Repository) DeleteCountries(country *models.Countries) error {
	query := db.Data.Table("countries").Where("id =?", country.ID).Delete(country)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r Repository) CountryStatus(country *models.Countries) error {
	tx := db.Data.Where("id", country.ID).Table("countries").Scan(&country)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if country.Active == true {
		country.Active = false
	} else {
		country.Active = true
	}
	tx = db.Data.Where("id", country.ID).Table("countries").Update("active", country.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) TotalPageCountry(limit int64) (int64, error) {
	var length int64
	if limit == 0 {
		limit = 10
	}

	query := db.Data.Table("countries").Count(&length)
	if query.Error != nil {
		log.Println(query.Error, "error in TotalPageCountry")
		return 0, query.Error
	}

	totalPage := length / limit
	if length%limit != 0 {
		totalPage++
	}
	log.Println(totalPage, "test")
	return totalPage, nil
}
