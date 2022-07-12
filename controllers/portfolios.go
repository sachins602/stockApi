package controllers

import (
	"net/http"

	md "goapi/middlewares"
	"goapi/models"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

func GetPortfolios(c *gin.Context) {

	var portfolios []models.Portfolio

	page := com.StrTo(c.Query("page")).MustInt()
	if page < 1 {
		page = 1
	}

	limit := com.StrTo(c.Query("limit")).MustInt()
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	models.DB.Limit(limit).Offset(offset).Order("id desc").Find(&portfolios)

	var total_data int64
	models.DB.Table("portfolios").Count(&total_data)

	c.JSON(http.StatusOK, gin.H{"page": page, "limit": limit, "total": total_data, "data": portfolios})

}

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
	models.DB.Create(&portfolio)

	c.JSON(http.StatusOK, gin.H{"data": portfolio})

}

func GetPortfolioByID(c *gin.Context) {
	var portfolio []models.Portfolio

	if err := models.DB.Where("username = ?", c.Param("username")).Find(&portfolio).Error; err != nil {
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
	if err := models.DB.Where("username = ?", c.Param("username")).Find(&portfolio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdatePortfolioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	if err := models.DB.Where("username = ? AND scrip = ?", c.Param("username"), c.Param("scrip")).Delete(&portfolio).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})

}
