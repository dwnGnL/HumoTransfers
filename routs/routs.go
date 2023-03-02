package routs

import (
	"Humo/handlers"

	"github.com/gin-gonic/gin"
)

type Routs struct {
	Engine   *gin.Engine
	Handlers *handlers.Handler
}

func NewHandler(engine *gin.Engine, handlers *handlers.Handler) *Routs {
	return &Routs{
		Engine:   engine,
		Handlers: handlers,
	}
}

func (r *Routs) Init() {
	//r.Engine.GET("/migrate", r.Handlers.Migration) // для миграции таблиц, вызывается один раз
	r.Engine.GET("/add_country", r.Handlers.AddCountry)
	r.Engine.GET("/add_language", r.Handlers.AddLanguage)
	r.Engine.GET("/add_sys_message", r.Handlers.AddSysMessage)
	r.Engine.GET("/add_currency", r.Handlers.AddCurrency)
	r.Engine.GET("/add_test", r.Handlers.AddTest)
	r.Engine.GET("/add_agents", r.Handlers.AddAgent)
	r.Engine.GET("/add_acc_agents", r.Handlers.AddAccount)

	r.Engine.GET("/get_country", r.Handlers.GetCountry)
	r.Engine.GET("/get_language", r.Handlers.GetLanguage)
	r.Engine.GET("/get_sys_message", r.Handlers.GetSysMessage)
	r.Engine.GET("/get_currency", r.Handlers.GetCurrency)
	r.Engine.GET("/get_test", r.Handlers.GetTest)
	r.Engine.GET("/get_agents", r.Handlers.GetAgent)
	r.Engine.GET("/get_acc_agents", r.Handlers.GetAccountAgent)

	r.Engine.GET("/update_countries", r.Handlers.UpdateCountries)
	r.Engine.GET("/update_language", r.Handlers.UpdateLanguage)
	r.Engine.GET("/update_sys_message", r.Handlers.UpdateSysMessage)
	r.Engine.GET("/update_currency", r.Handlers.UpdateCurrency)
	r.Engine.GET("/update_test", r.Handlers.UpdateTest)
	//r.Engine.GET("/update_agent:", r.Handlers.UpdateAgents)
	r.Engine.GET("/update_account", r.Handlers.UpdateAccountAgent)
	r.Engine.GET("/update_account_def", r.Handlers.UpdateAccountDefault)

	r.Engine.GET("/status_countries", r.Handlers.CountryStatus)
	r.Engine.GET("/status_language", r.Handlers.LanguageStatus)
	r.Engine.GET("/status_sys_message", r.Handlers.SysMessageStatus)
	r.Engine.GET("/status_agent", r.Handlers.AgentStatus)
	r.Engine.GET("/status_account", r.Handlers.AccountAgentStatus)
}

//http://localhost:8080/add_language
//http://localhost:8080/add_sys_message
//http://localhost:8080/add_currency
//http://localhost:8080/add_test
//http://localhost:8080/add_agents
//http://localhost:8080/add_acc_agents
//
//http://localhost:8080/get_country
//http://localhost:8080/get_language
//http://localhost:8080/get_sys_message
//http://localhost:8080/get_currency
//http://localhost:8080/get_test
//http://localhost:8080/get_agents
//http://localhost:8080/get_acc_agents
//
//http://localhost:8080/update_countries
//http://localhost:8080/update_language
//http://localhost:8080/update_sys_message
//http://localhost:8080/update_currency
//http://localhost:8080/update_test
//http://localhost:8080/update_agent
//http://localhost:8080/update_account
//http://localhost:8080/update_account_def
//
//
//http://localhost:8080/status_countries
//http://localhost:8080/status_language
//http://localhost:8080/status_sys_message
//http://localhost:8080/status_agent
//http://localhost:8080/status_account

// todo ngrok http 8080
