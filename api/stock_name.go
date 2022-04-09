package api

import (
	// "errors"
	"goapi/database"
	"goapi/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRepo struct {
	Db *gorm.DB
}

func New() *UserRepo {
	db := database.InitDb()
	// db.AutoMigrate(&models.Stock_details{})
	return &UserRepo{Db: db}
}

// //create user
// func (repository *UserRepo) CreateUser(c *gin.Context) {
//    var user models.User
//    c.BindJSON(&user)
//    err := models.CreateUser(repository.Db, &user)
//    if err != nil {
//       c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//       return
//    }
//    c.JSON(http.StatusOK, user)
// }

// get Stocks
func (repository *UserRepo) GetStocks(c *gin.Context) {
	var stocks []models.Stock_details
	err := models.GetStocks(repository.Db, &stocks)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, stocks)

}

//get Broker
func (repository *UserRepo) GetBroker(c *gin.Context) {
	var broker []models.Broker_details
	err := models.GetBroker(repository.Db, &broker)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, broker)

}

func (repository *UserRepo) GetSector(c *gin.Context) {
	var sector []models.Sector_details
	err := models.GetSector(repository.Db, &sector)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, sector)

}

func (repository *UserRepo) GetIndex(c *gin.Context) {
	var index []models.Index_details
	err := models.GetIndex(repository.Db, &index)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, index)

}

// //get user by id
// func (repository *UserRepo) GetUser(c *gin.Context) {
//    id, _ := c.Params.Get("id")
//    var user models.User
//    err := models.GetUser(repository.Db, &user, id)
//    if err != nil {
//       if errors.Is(err, gorm.ErrRecordNotFound) {
//          c.AbortWithStatus(http.StatusNotFound)
//          return
//       }

//       c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//       return
//    }
//    c.JSON(http.StatusOK, user)
// }

// // update user
// func (repository *UserRepo) UpdateUser(c *gin.Context) {
//    var user models.User
//    id, _ := c.Params.Get("id")
//    err := models.GetUser(repository.Db, &user, id)
//    if err != nil {
//       if errors.Is(err, gorm.ErrRecordNotFound) {
//          c.AbortWithStatus(http.StatusNotFound)
//          return
//       }

//       c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//       return
//    }
//    c.BindJSON(&user)
//    err = models.UpdateUser(repository.Db, &user)
//    if err != nil {
//       c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//       return
//    }
//    c.JSON(http.StatusOK, user)
// }

// // delete user
// func (repository *UserRepo) DeleteUser(c *gin.Context) {
//    var user models.User
//    id, _ := c.Params.Get("id")
//    err := models.DeleteUser(repository.Db, &user, id)
//    if err != nil {
//       c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
//       return
//    }
//    c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
// }
