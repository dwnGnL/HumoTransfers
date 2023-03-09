package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

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

func (h *Handler) GetLanguage(ctx *gin.Context) {
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")

	if pageStr == "" {
		pageStr = "0"
	}
	if limitStr == "" {
		limitStr = "0"
	}
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetLanguages(offset, limit)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	totalPage, err := h.Repository.TotalPageLanguage(int64(limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var Language models.LanguageWithPage
	Language.Languages = app
	Language.TotalPage = totalPage

	ctx.JSON(http.StatusOK, Language)
}

func (h Handler) UpdateLanguage(ctx *gin.Context) {
	var language *models.Languages

	if err := ctx.ShouldBindJSON(&language); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if language.Name == "" && language.Icon == "" {
		err := h.Repository.DeleteLanguage(language)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateLanguage(language)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
