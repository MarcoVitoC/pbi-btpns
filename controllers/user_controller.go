package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/MarcoVitoC/pbi-btpns/database"
	"github.com/MarcoVitoC/pbi-btpns/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	user := &models.User{}

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Request failed!",
		})

		return
	}

	var validate = validator.New()
	if validationErr := validate.Struct(user); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": validationErr.Error(),
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Failed to hash password!",
		})

		return
	}

	newUser := models.User{Username: user.Username, Email: user.Email, Password: string(hash)}
	result := database.DatabaseConnection().Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": result.Error,
		})

		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Registration succeed!",
	})
}

func Login(c *gin.Context) {
	user := &models.User{}
	loginRequest := &models.LoginRequest{}

	if err := c.Bind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Request failed!",
		})

		return
	}

	db := database.DatabaseConnection()
	db.First(&user, "email = ?", loginRequest.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "User not found!",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Invalid password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"message": "Failed to create token!",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Login succeed!",
	})
}

func Update(c *gin.Context) {
	//
}

func Delete(c *gin.Context) {
	//
}