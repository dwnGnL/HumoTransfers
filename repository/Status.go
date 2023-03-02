package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

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

func (r Repository) AgentStatus(agent *models.Agents) error {
	tx := db.Data.Where("id", agent.ID).Table("agents").Scan(&agent)
	if tx.Error != nil {
		log.Println(tx.Error)
		return tx.Error
	}
	if agent.Active == true {
		agent.Active = false
	} else {
		agent.Active = true
	}
	tx = db.Data.Where("id", agent.ID).Table("agents").Update("active", agent.Active)
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
