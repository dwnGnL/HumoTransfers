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

func (r *Repository) GetCountries(offset int, limit int) ([]*models.Countries, int64, error) {
	var countries []*models.Countries
	var length int64
	if limit == 0 {
		limit = 10
	}

	query := db.Data.Table("countries")
	//if filter!="" {
	//	query = query.Where("name like ?","%"+filter+"%")
	//}
	tx := query.Count(&length)
	tx = query.Limit(limit).Offset(offset).Find(&countries)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, 0, tx.Error
	}
	return countries, length, nil
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
