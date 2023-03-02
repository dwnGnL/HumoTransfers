package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h Handler) CountryStatus(ctx *gin.Context) {
	var country *models.Countries

	if err := ctx.ShouldBindJSON(&country); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.CountryStatus(country)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
func (h Handler) LanguageStatus(ctx *gin.Context) {
	var language *models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.LanguageStatus(language)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}

func (h Handler) SysMessageStatus(ctx *gin.Context) {
	var message *models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.SysMessageStatus(message)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}

func (h Handler) AgentStatus(ctx *gin.Context) {
	var agent *models.Agents

	if err := ctx.ShouldBindJSON(&agent); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AgentStatus(agent)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}

func (h Handler) AccountAgentStatus(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AccountAgentStatus(account)
	if err != nil {
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
