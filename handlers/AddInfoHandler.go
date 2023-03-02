package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddCountry(ctx *gin.Context) {
	var country models.Countries

	if err := ctx.ShouldBindJSON(&country); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	country.Active = true
	err := h.Repository.AddCountry(&country)
	if err != nil {
		log.Printf("%s in AddCountry", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Country added!")
}

func (h *Handler) AddLanguage(ctx *gin.Context) {
	var language models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	language.Active = true
	err := h.Repository.AddLanguage(&language)
	if err != nil {
		log.Printf("%s in AddLanguage(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Language added!")
}

func (h *Handler) AddSysMessage(ctx *gin.Context) {
	var message models.SysMessage

	if err := ctx.ShouldBindJSON(&message); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	message.Active = true
	err := h.Repository.AddSysMessage(&message)
	if err != nil {
		log.Printf("%s in AddSysMessage(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "SysMessage added!")
}

func (h *Handler) AddCurrency(ctx *gin.Context) {
	var currency models.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AddCurrency(&currency)
	if err != nil {
		log.Printf("%s in AddCurrency(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Currency added!")
}

func (h *Handler) AddTest(ctx *gin.Context) {
	var test models.Test

	if err := ctx.ShouldBindJSON(&test); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AddTest(&test)
	if err != nil {
		log.Printf("%s in AddTest(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Test added!")
}

func (h *Handler) AddAgent(ctx *gin.Context) {
	var agents models.Agents

	if err := ctx.ShouldBindJSON(&agents); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	agents.Active = true
	err := h.Repository.AddAgent(&agents)
	if err != nil {
		log.Printf("%s in AddAgent(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Agent added!")
}

func (h *Handler) AddAccount(ctx *gin.Context) {
	var account models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	account.Active = true
	err := h.Repository.AddAccount(&account)
	if err != nil {
		log.Printf("%s in AddAccount(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "AccountAgent added!")
}
