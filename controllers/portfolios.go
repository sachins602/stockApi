package controllers

import (
	"fmt"
	"net/http"

	md "goapi/middlewares"
	"goapi/models"

	"github.com/gin-gonic/gin"
)

type CreatePortfolioInput struct {
	Scrip string  `json:"scrip" binding:"required"`
	Total float64 `json:"total" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

func CreatePortfolio(c *gin.Context) {
	var input CreatePortfolioInput

	currentUser := md.GetCurrentUser()

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.Portfolio{}})
		return
	}

	portfolio := models.Portfolio{Username: currentUser, Scrip: input.Scrip, Total: input.Total, Price: input.Price}

	models.DB.Where("username = ? AND scrip = ?", currentUser, input.Scrip).FirstOrCreate(&portfolio)
	c.JSON(http.StatusOK, gin.H{"data": portfolio})

}

func GetPortfolioByID(c *gin.Context) {
	var portfolio []models.PortfolioResponseForTable

	query := fmt.Sprintf("SELECT portfolios.*, stocks.LastPrice, stocks.Open, (portfolios.total * stocks.LastPrice) - (portfolios.total * portfolios.price) as TotalProfit FROM portfolios LEFT JOIN stocks ON portfolios.scrip = stocks.StockName WHERE portfolios.username = '%s' ORDER BY portfolios.created_at;", c.Param("username"))

	if err := models.DB.Raw(query).Scan(&portfolio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, portfolio)
}

//schema validation
type UpdatePortfolioInput struct {
	Scrip string  `json:"scrip"`
	Total float64 `json:"total"`
	Price float64 `json:"price"`
}

// Update a portfolio
func UpdatePortfolio(c *gin.Context) {
	// Get model if exist
	var portfolio models.Portfolio

	var currentUser = md.GetCurrentUser()
	// Validate input
	var input UpdatePortfolioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Where("username = ? AND scrip = ?", currentUser, input.Scrip).First(&portfolio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Model(&portfolio).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": portfolio})
}

// Delete a Portfolio
func DeletePortfolio(c *gin.Context) {
	// Get model if exist
	var portfolio models.Portfolio
	//db.Where("name = ? AND age >= ?", "jinzhu", "22").Find(&users)
	if err := models.DB.Unscoped().Where("username = ? AND scrip = ?", c.Param("username"), c.Param("scrip")).Delete(&portfolio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})

}
