package controllers

import (
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get Stocks
func GetStocks(c *gin.Context) {

	var stocks []models.Stock

	models.DB.Find(&stocks)

	c.JSON(http.StatusOK, stocks)

}

//get Broker
func GetBroker(c *gin.Context) {
	var brokers []models.Broker

	models.DB.Find(&brokers)

	c.JSON(http.StatusOK, brokers)

}

// get sector
func GetSector(c *gin.Context) {
	var sectors []models.Sector

	models.DB.Find(&sectors)

	c.JSON(http.StatusOK, sectors)

}

//get index
func GetIndex(c *gin.Context) {
	var indices []models.Index

	models.DB.Find(&indices)

	c.JSON(http.StatusOK, indices)
}

//get gainer
func GetGainer(c *gin.Context) {
	var gainers []models.Gainer

	models.DB.Table("top_gainers").Find(&gainers)

	c.JSON(http.StatusOK, gainers)
}

//get loser
func GetLoser(c *gin.Context) {
	var losers []models.Gainer

	models.DB.Table("top_losers").Find(&losers)

	c.JSON(http.StatusOK, losers)
}

//get subIndex
func GetSubIndex(c *gin.Context) {
	var subIndices []models.SubIndex

	models.DB.Table("sub_indices").Find(&subIndices)

	c.JSON(http.StatusOK, subIndices)
}

//get NEPSE data
func GetNepse(c *gin.Context) {
	var nepse models.Index

	models.DB.First(&nepse, "IndexName = ?", "NEPSE")

	c.JSON(http.StatusOK, nepse)
}

// get NEPSE index history data
func GetNepseIndexHistory(c *gin.Context) {
	var nepseHistoric []models.Historic

	if err := models.DB.Where("time > 1635865190").Table("historic").Find(&nepseHistoric).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.JSON(http.StatusOK, nepseHistoric)
}

//get NEPSE history data
func GetNepseHistory(c *gin.Context) {
	var historics []models.Historic

	//models.DB.Find(&historics)

	if err := models.DB.Where("scrip = ? AND time > 1622732399", c.Param("scrip")).Table(c.Param("sector")).Find(&historics).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// var total_data int64
	// models.DB.Where("scrip = ?", c.Param("scrip")).Table(c.Param("sector")).Count(&total_data)

	c.JSON(http.StatusOK, historics)
}
