package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) AddAccount(account *models.AccountAgent) error {
	tx := db.Data.Create(account)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
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

func (r Repository) UpdateAccountAgent(AccAgent *models.AccountAgent) error {

	var AccAgents *models.AccountAgent
	query := db.Data.Where("id=?", AccAgent.ID).Find(&AccAgents)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}

	tx := db.Data.Model(&models.AccountAgent{}).Where("id = ?", AccAgent.ID).Updates(models.AccountAgent{AgentId: AccAgent.AgentId, CurrencyId: AccAgent.CurrencyId, IsDefault: AccAgent.IsDefault, Type: AccAgent.Type})
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}

	return nil
}

func (r Repository) DeleteAccountAgent(AccAgent *models.AccountAgent) error {

	query := db.Data.Table("account_agents").Where("id =?", AccAgent.ID).Delete(AccAgent)
	if query.Error != nil {
		log.Println(query.Error)
		return query.Error
	}
	return nil
}

func (r Repository) UpdateAccountDefault(AccAgent *models.AccountAgent) error {
	tx := db.Data.Where("id", AccAgent.ID).Table("account_agents").Scan(&AccAgent)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if AccAgent.IsDefault == true {
		AccAgent.IsDefault = false
	} else {
		AccAgent.IsDefault = true
	}
	tx = db.Data.Where("id", AccAgent.ID).Table("account_agents").Update("is_default", AccAgent.IsDefault)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}

func (r Repository) AccountAgentStatus(AccAgent *models.AccountAgent) error {
	tx := db.Data.Where("id", AccAgent.ID).Table("account_agents").Scan(&AccAgent)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if AccAgent.Active == true {
		AccAgent.Active = false
	} else {
		AccAgent.Active = true
	}
	tx = db.Data.Where("id", AccAgent.ID).Table("account_agents").Update("active", AccAgent.Active)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	return nil
}
