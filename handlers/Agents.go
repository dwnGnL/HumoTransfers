package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

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

func (h *Handler) GetAgent(ctx *gin.Context) {

	pagination := GeneratePaginationFromRequest(ctx)
	AgentLists, err := h.Repository.GetAgent(&pagination)
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
	pagination.Records = AgentLists
	pagination.TotalPages = TotalPages
	ctx.JSON(http.StatusOK, pagination)
}

func (h Handler) UpdateAgents(ctx *gin.Context) {
	var agents *models.Agents
	ctx.Param("id")
	if err := ctx.ShouldBindJSON(&agents); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if agents.Name == "" && agents.LegalName == "" {
		err := h.Repository.DeleteAgents(agents)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateAgents(agents)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}
	ctx.JSON(http.StatusOK, " Done!")
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
