package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

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

	//switch  {
	//case filter == "":
	//	return countries, nil
	//
	//}
	return countries, length, nil
}

func (r *Repository) GetLanguages(offset int, limit int) ([]*models.Languages, error) {
	var language []*models.Languages
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("language").Limit(limit).Offset(offset).Find(&language)

	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return language, nil
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

func (r *Repository) GetTest(offset int, limit int) ([]*models.Test, error) {
	var test []*models.Test
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("test").Limit(limit).Offset(offset).Find(&test)

	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return test, nil
}

func (r *Repository) GetAgent(offset int, limit int) ([]*models.Agents, error) {
	var agents []*models.Agents
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("agents").Limit(limit).Offset(offset).Find(&agents)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return agents, nil
}

func (r *Repository) GetAccountAgent(offset int, limit int) ([]*models.AccountAgent, error) {
	var AccAgents []*models.AccountAgent
	if limit == 0 {
		limit = 10
	}
	tx := db.Data.Table("account_agents").Limit(limit).Offset(offset).Find(&AccAgents)
	if tx.Error != nil {
		log.Println(tx.Error)
		return nil, tx.Error
	}
	return AccAgents, nil
}
