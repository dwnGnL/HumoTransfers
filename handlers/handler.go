package handlers

import (
	"Humo/repository"
)

type Handler struct {
	Repository *repository.Repository
}

func NewHandler(repository *repository.Repository) *Handler {
	return &Handler{
		Repository: repository,
	}
}

//func (h *Handler) Migration(ctx *gin.Context) {
//
//	err := h.Repository.Migrate()
//	if err != nil {
//		log.Printf("%s in AddUser(server)", err)
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
//		return
//	}
//	ctx.JSON(http.StatusOK, "Migration done!")
//}
