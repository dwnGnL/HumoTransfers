package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetAccountAgent(offset, limit)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	totalPage, err := h.Repository.TotalPageAgents(int64(limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var Account models.AccountWithPage
	Account.Account = app
	Account.TotalPage = totalPage

	ctx.JSON(http.StatusOK, Account)
}

func (h Handler) UpdateAccountAgent(ctx *gin.Context) {
	var account *models.AccountAgent

	if err := ctx.ShouldBindJSON(&account); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if account.AgentId == 0 && account.CurrencyId == 0 &&
		account.Type == "" {
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
