package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddServCountry(ctx *gin.Context) {
	var servicesCountry models.ServicesCountry

	if err := ctx.ShouldBindJSON(&servicesCountry); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	servicesCountry.Active = true
	err := h.Repository.AddServCountry(&servicesCountry)
	if err != nil {
		log.Printf("%s in AddServCountry(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "AddServCountry added!")
}

func (h *Handler) GetServCountry(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)

	ServCountryLists, err := h.Repository.GetServCountry(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageServCountry(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = ServCountryLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) DeleteServCountry(ctx *gin.Context) {
	var servCountry *models.ServicesCountry

	if err := ctx.ShouldBindJSON(&servCountry); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.DeleteServCountry(servCountry)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) ServCountryStatus(ctx *gin.Context) {
	var servicesCountry *models.ServicesCountry

	if err := ctx.ShouldBindJSON(&servicesCountry); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.ServCountryStatus(servicesCountry)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
