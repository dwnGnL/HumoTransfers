package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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

func (h *Handler) GetCurrency(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetCurrency(offset, limit)
	if err != nil {
		log.Println(err)
		return
	}
	totalPage, err := h.Repository.TotalPageCurrency(int64(limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var Currency models.CurrencyWithPage
	Currency.Currency = app
	Currency.TotalPage = totalPage

	ctx.JSON(http.StatusOK, Currency)
}

func (h Handler) UpdateCurrency(ctx *gin.Context) {
	var currency *models.Currency

	if err := ctx.ShouldBindJSON(&currency); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if currency.Name == "" && currency.Icon == "" {
		err := h.Repository.DeleteCurrency(currency)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateCurrency(currency)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
