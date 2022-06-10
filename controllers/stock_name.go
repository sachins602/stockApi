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

//get NEPSE history data
func GetNepseHistory(c *gin.Context) {
	var historics []models.Historic

	models.DB.Find(&historics)

	var total_data int64
	models.DB.Table("historics").Count(&total_data)

	c.JSON(http.StatusOK, gin.H{"total": total_data, "data": historics})
}
