package handlers

import (
	"Humo/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) GetCountry(ctx *gin.Context) {
	//pageStr := ctx.Param("page")
	//limitStr := ctx.Param("limit")

	queries := ctx.Request.URL.Query()
	pageStr := queries.Get("page")
	limitStr := queries.Get("limit")
	log.Println(pageStr, limitStr)
	if pageStr == "" {
		pageStr = "0"
	}
	if limitStr == "" {
		limitStr = "0"
	}

	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitStr)
	//log.Println(limit, page)
	offset := (page - 1) * 10
	app, err := h.Repository.GetCountries(offset, limit)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	totalPage, err := h.Repository.TotalPageCountry(int64(limit))
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	var Country models.CountriesWithPage
	Country.Countries = app
	Country.TotalPage = totalPage
	log.Println(totalPage, "handler")
	ctx.JSON(http.StatusOK, Country)
}

func (h Handler) UpdateCountries(ctx *gin.Context) {
	var countries *models.Countries

	if err := ctx.ShouldBindJSON(&countries); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if countries.Icon == "" && countries.Name == "" {
		err := h.Repository.DeleteCountries(countries)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	} else {
		err := h.Repository.UpdateCountries(countries)
		log.Println("work&")
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			log.Println(err)
			return
		}
	}

	ctx.JSON(http.StatusOK, " Done!")
}

func (h Handler) CountryStatus(ctx *gin.Context) {
	var country *models.Countries

	if err := ctx.ShouldBindJSON(&country); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	err := h.Repository.CountryStatus(country)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		log.Println(err)
		return
	}
	ctx.JSON(http.StatusOK, "Done!")
}
