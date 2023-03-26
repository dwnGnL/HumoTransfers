package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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

func (h *Handler) GetAccountAgent(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)

	AccLists, err := h.Repository.GetAccountAgent(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageAgents(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = AccLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdateAccountAgent(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if account.AgentId == 0 && account.CurrencyId == 0 &&
		account.Type == 0 {
		err := h.Repository.DeleteAccountAgent(account)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	} else {
		err := h.Repository.UpdateAccountAgent(account)
		if err != nil {
			log.Println(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, " Done!")
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
