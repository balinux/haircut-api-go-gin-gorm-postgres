package controllers

import (
	"haircut-api-go-gin-gorm-postgres/models"
	"haircut-api-go-gin-gorm-postgres/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HaircutController struct {
	service services.HaircutService
}

func NewHaircutController(service services.HaircutService) *HaircutController {
	return &HaircutController{service}
}

func (ctrl *HaircutController) GetAllHaircuts(c *gin.Context) {
	haircuts, err := ctrl.service.GetAllHaircuts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": haircuts})
}

func (c *HaircutController) CreateHaircut(ctx *gin.Context) {
	var haircut models.Haircut
	if err := ctx.ShouldBindJSON(&haircut); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	createHaircut, err := c.service.CreateHaircut(haircut)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusCreated, createHaircut)
}

func (ctrl *HaircutController) GetHaircut(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid Id"})
		return
	}
	haircut, err := ctrl.service.GetHaircutByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "haircut not found"})
		return
	}

	c.JSON(http.StatusOK, haircut)
}
