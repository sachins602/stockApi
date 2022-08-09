package controllers

import (
	// "encoding/json"
	// "io/ioutil"
	"errors"
	"fmt"
	"net/http"

	"goapi/auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	// "github.com/victorsteven/fullstack/api/models"
	// "github.com/victorsteven/fullstack/api/responses"
	// "github.com/victorsteven/fullstack/api/utils/formaterror"
	md "goapi/middlewares"
	"goapi/models"
	"goapi/utils"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utils.Simple(verr)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Password: input.Password}
	user.Prepare()
	token, err := SignIn(user.Username, user.Password)
	fmt.Println(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	md.SaveCurrentUser(user.Username)

	c.JSON(http.StatusOK, gin.H{"token": token, "username": user.Username})

	// body, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// user := models.User{}
	// err = json.Unmarshal(body, &user)
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }

	// user.Prepare()
	// err = user.Validate("login")
	// if err != nil {
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, err)
	// 	return
	// }
	// token, err := server.SignIn(user.Email, user.Password)
	// if err != nil {
	// 	formattedError := formaterror.FormatError(err.Error())
	// 	responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
	// 	return
	// }
	// responses.JSON(w, http.StatusOK, token)
}

func SignIn(username, password string) (string, error) {

	var err error

	user := models.User{}

	err = models.DB.Model(models.User{}).Where("username = ?", username).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	return auth.CreateToken(user.ID)
}

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput
	var err error

	if err := c.ShouldBindJSON(&input); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utils.Simple(verr)})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Username: input.Username, Fullname: input.Fullname, Email: input.Email, Password: input.Password}

	user.Prepare()

	// err = user.BeforeSave()

	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	fmt.Println(user.Username, user.Password)
	// return
	res, err := user.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success", "data": res})

}

//get userdetails by username
func GetUserByUsername(c *gin.Context) {
	username := c.Param("username")
	user := models.User{}
	err := models.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
