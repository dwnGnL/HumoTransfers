package repository

import (
	"Humo/db"
	"Humo/models"
	"log"
)

func (r Repository) UpdateCountries(country *models.Countries) error {
	if country.Name == "" && country.Icon == "" {
		query := db.Data.Table("countries").Where("id =?", country.ID).Delete(country)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {
		var countries *models.Countries
		query2 := db.Data.Where("id=?", country.ID).Find(&countries)
		if query2.Error != nil {
			log.Println(query2.Error)
			return query2.Error
		}
		if country.Name == "" {
			country.Name = countries.Name
		}
		if country.Icon == "" {
			country.Icon = countries.Icon
		}
		log.Println(1, country, 2, countries)
		//tx := db.Data.Model(&country).Where("id = ?", country.ID).Updates(models.Countries{Name: country.Name, Icon: country.Icon})
		tx := db.Data.Model(&models.Countries{}).Where("id = ?", country.ID).Updates(models.Countries{Name: country.Name, Icon: country.Icon, Active: country.Active})
		//tx := db.Data.Table("countries").Save(&country)
		if tx.Error != nil {
			log.Println(tx.Error)
			return tx.Error
		}
	}
	return nil
}

func (r Repository) UpdateLanguage(language *models.Languages) error {
	if language.Name == "" && language.Icon == "" {
		query := db.Data.Table("languages").Where("id =?", language.ID).Delete(language)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {
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
	}
	return nil
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

func (r Repository) UpdateCurrency(currency *models.Currency) error {
	if currency.Name == "" && currency.Icon == "" {
		query := db.Data.Table("currencies").Where("id =?", currency.ID).Delete(currency)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {
		var currencies *models.Languages
		query := db.Data.Where("id=?", currency.ID).Find(&currencies)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
		if currency.Name == "" {
			currency.Name = currencies.Name
		}
		if currency.Icon == "" {
			currency.Icon = currencies.Icon
		}

		tx := db.Data.Model(&models.Currency{}).Where("id = ?", currency.ID).Updates(models.Currency{Name: currency.Name, Icon: currency.Icon})
		if tx.Error != nil {
			log.Println(tx.Error)
			return tx.Error
		}
	}
	return nil
}

// todo need id
func (r Repository) UpdateTest(testTable *models.Test) error {
	log.Println(testTable.ID)
	if testTable.Entity == "" && testTable.EntityId == 0 && testTable.LangId == 0 && testTable.Value == "" {
		query := db.Data.Table("test").Where("id =?", testTable.ID).Delete(testTable)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {
		var tests *models.Test
		query := db.Data.Table("test").Where("id=?", testTable.ID).Find(&tests)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
		if testTable.Entity == "" {
			testTable.Entity = tests.Entity
		}
		if testTable.EntityId == 0 {
			testTable.EntityId = tests.EntityId
		}
		if testTable.LangId == 0 {
			testTable.LangId = tests.LangId
		}
		if testTable.Value == "" {
			testTable.Value = tests.Value
		}

		tx := db.Data.Table("test").Model(&models.Test{}).Where("id = ?", testTable.ID).Updates(models.Test{Entity: testTable.Entity, EntityId: testTable.EntityId, LangId: testTable.LangId, Value: testTable.Value})
		if tx.Error != nil {
			log.Println(tx.Error)
			return tx.Error
		}
	}
	return nil
}

func (r Repository) UpdateAgents(agent *models.Agents) error {
	if agent.Name == "" && agent.LegalName == "" {
		query := db.Data.Table("agents").Where("id =?", agent.ID).Delete(agent)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {
		var agents *models.Agents
		query := db.Data.Where("id=?", agent.ID).Find(&agents)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
		if agent.Name == "" {
			agent.Name = agents.Name
		}
		if agent.LegalName == "" {
			agent.LegalName = agents.LegalName
		}

		tx := db.Data.Model(&models.Agents{}).Where("id = ?", agent.ID).Updates(models.Agents{Name: agent.Name, LegalName: agent.LegalName})
		if tx.Error != nil {
			log.Println(tx.Error)
			return tx.Error
		}
	}
	return nil
}

func (r Repository) UpdateAccountAgent(AccAgent *models.AccountAgent) error {
	log.Println(AccAgent)
	if AccAgent.AgentId == 0 && AccAgent.CurrencyId == 0 && AccAgent.Type == "" {
		query := db.Data.Table("account_agents").Where("id =?", AccAgent.ID).Delete(AccAgent)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
	} else {
		var AccAgents *models.AccountAgent
		query := db.Data.Where("id=?", AccAgent.ID).Find(&AccAgents)
		if query.Error != nil {
			log.Println(query.Error)
			return query.Error
		}
		if AccAgent.AgentId == 0 {
			AccAgent.AgentId = AccAgents.AgentId
		}
		if AccAgent.CurrencyId == 0 {
			AccAgent.CurrencyId = AccAgents.CurrencyId
		}
		if AccAgent.Type == "" {
			AccAgent.Type = AccAgents.Type
		}

		tx := db.Data.Model(&models.AccountAgent{}).Where("id = ?", AccAgent.ID).Updates(models.AccountAgent{AgentId: AccAgent.AgentId, CurrencyId: AccAgent.CurrencyId, IsDefault: AccAgent.IsDefault, Type: AccAgent.Type})
		if tx.Error != nil {
			log.Println(tx.Error)
			return tx.Error
		}
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
