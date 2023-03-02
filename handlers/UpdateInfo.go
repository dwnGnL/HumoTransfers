package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h Handler) UpdateCountries(ctx *gin.Context) {
	var countries *models.Countries

	if err := ctx.ShouldBindJSON(&countries); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateCountries(countries)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateLanguage(ctx *gin.Context) {
	var language *models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateLanguage(language)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateSysMessage(ctx *gin.Context) {
	var message *models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateSysMessage(message)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateCurrency(ctx *gin.Context) {
	var currency *models.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateCurrency(currency)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateTest(ctx *gin.Context) {
	var testTable *models.Test

	if err := ctx.ShouldBindJSON(&testTable); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateTest(testTable)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateAgents(ctx *gin.Context) {
	var agents *models.Agents
	ctx.Param("id")
	if err := ctx.ShouldBindJSON(&agents); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateAgents(agents)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateAccountAgent(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateAccountAgent(account)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) UpdateAccountDefault(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := h.Repository.UpdateAccountDefault(account)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
}
