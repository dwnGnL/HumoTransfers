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
	r.Engine.POST("/add_country", r.Handlers.AddCountry)
	r.Engine.POST("/add_language", r.Handlers.AddLanguage)
	r.Engine.POST("/add_sys_message", r.Handlers.AddSysMessage)
	r.Engine.POST("/add_currency", r.Handlers.AddCurrency)
	r.Engine.POST("/add_transfer", r.Handlers.AddTransfer)
	r.Engine.POST("/add_agents", r.Handlers.AddAgent)
	r.Engine.POST("/add_acc_agents", r.Handlers.AddAccount)
	r.Engine.POST("/add_user", r.Handlers.AddUserInfo)
	r.Engine.POST("/add_vendor", r.Handlers.AddVendor)
	r.Engine.POST("/add_service", r.Handlers.AddServices)
	r.Engine.POST("/add_serv_country", r.Handlers.AddServCountry)
	r.Engine.POST("/add_serv_rules", r.Handlers.AddServiceRules)
	r.Engine.POST("/add_payment_type", r.Handlers.AddPaymentType)

	r.Engine.GET("/get_country", r.Handlers.GetCountry)
	r.Engine.GET("/get_language", r.Handlers.GetLanguage)
	r.Engine.GET("/get_sys_message", r.Handlers.GetSysMessage)
	r.Engine.GET("/get_currency", r.Handlers.GetCurrency)
	r.Engine.GET("/get_transfer", r.Handlers.GetTransfer)
	r.Engine.GET("/get_agents", r.Handlers.GetAgent)
	r.Engine.GET("/get_acc_agents", r.Handlers.GetAccountAgent)
	r.Engine.GET("/get_user", r.Handlers.GetUserInfo)
	r.Engine.GET("/get_vendor", r.Handlers.GetVendor)
	r.Engine.GET("/get_services", r.Handlers.GetServices)
	r.Engine.GET("/get_serv_country", r.Handlers.GetServCountry)
	r.Engine.GET("/get_serv_rules", r.Handlers.GetServiceRules)
	r.Engine.GET("/get_payment_type", r.Handlers.GetPaymentType)

	r.Engine.PUT("/update_countries", r.Handlers.UpdateCountries)
	r.Engine.PUT("/update_language", r.Handlers.UpdateLanguage)
	r.Engine.PUT("/update_sys_message", r.Handlers.UpdateSysMessage)
	r.Engine.PUT("/update_currency", r.Handlers.UpdateCurrency)
	r.Engine.PUT("/update_transfer", r.Handlers.UpdateTransfer)
	r.Engine.PUT("/update_agent", r.Handlers.UpdateAgents)
	r.Engine.PUT("/update_account", r.Handlers.UpdateAccountAgent)
	r.Engine.PUT("/update_account_def", r.Handlers.UpdateAccountDefault)
	r.Engine.PUT("/update_user", r.Handlers.UpdateUserInfo)
	r.Engine.PUT("/update_vendor", r.Handlers.UpdateVendor)
	r.Engine.PUT("/update_service", r.Handlers.UpdateService)
	r.Engine.PUT("/delete_serv_country", r.Handlers.DeleteServCountry)
	r.Engine.PUT("/update_serv_rules", r.Handlers.UpdateServiceRules)
	r.Engine.PUT("/update_payment_type", r.Handlers.UpdatePaymentType)

	r.Engine.PATCH("/status_countries", r.Handlers.CountryStatus)
	r.Engine.PATCH("/status_language", r.Handlers.LanguageStatus)
	r.Engine.PATCH("/status_sys_message", r.Handlers.SysMessageStatus)
	r.Engine.PATCH("/status_agent", r.Handlers.AgentStatus)
	r.Engine.PATCH("/status_account", r.Handlers.AccountAgentStatus)
	r.Engine.PATCH("/status_user", r.Handlers.UserInfoStatus)
	r.Engine.PATCH("/status_vendor", r.Handlers.VendorStatus)
	r.Engine.PATCH("/status_service", r.Handlers.ServiceStatus)
	r.Engine.PATCH("/status_serv_country", r.Handlers.ServCountryStatus)

	r.Engine.POST("/upload", r.Handlers.UploadImage)
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
