package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)

	offset := (page - 1) * 10
	app, err := h.Repository.GetAgent(offset, limit)
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
	var Agents models.AgentsWithPage
	Agents.Agents = app
	Agents.TotalPage = totalPage

	ctx.JSON(http.StatusOK, Agents)
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
