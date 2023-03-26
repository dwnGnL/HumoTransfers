package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) AddServiceRules(ctx *gin.Context) {
	var ServiceRules models.ServicesRules

	if err := ctx.ShouldBindJSON(&ServiceRules); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.AddServiceRules(&ServiceRules)
	if err != nil {
		log.Printf("%s in AddServiceRules(server)", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "ServiceRules added!")
}

func (h *Handler) GetServiceRules(ctx *gin.Context) {
	pagination := GeneratePaginationFromRequest(ctx)
	ServiceRulesList, err := h.Repository.GetServiceRules(&pagination)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	TotalPages, err := h.Repository.TotalPageServiceRules(int64(pagination.Limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	pagination.Records = ServiceRulesList
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdateServiceRules(ctx *gin.Context) {
	var ServiceRules *models.ServicesRules

	if err := ctx.ShouldBindJSON(&ServiceRules); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if ServiceRules.Name == "" && ServiceRules.Type == "" {
		err := h.Repository.DeleteServiceRules(ServiceRules)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateServiceRules(ServiceRules)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
}
