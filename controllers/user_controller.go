package controllers

import (
	"net/http"

	"github.com/MarcoVitoC/pbi-btpns/helpers"
	"github.com/MarcoVitoC/pbi-btpns/database"
	"github.com/MarcoVitoC/pbi-btpns/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	user := &models.User{}

	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Request failed!",
		})

		return
	}

	var validate = validator.New()
	if validationErr := validate.Struct(user); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": validationErr.Error(),
		})

		return
	}

	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Failed to hash password!",
		})

		return
	}

	newUser := models.User{Username: user.Username, Email: user.Email, Password: string(hash)}
	result := database.DB().Create(&newUser)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": result.Error,
		})

		return
	}
	
	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "Registration succeed!",
	})
}

func Login(c *gin.Context) {
	loginRequest := &models.LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Request failed!",
		})

		return
	}

	user := &models.User{}
	db := database.DB()
	db.First(&user, "email = ?", loginRequest.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "User not found!",
		})

		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Invalid password",
		})

		return
	}

	tokenString, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Failed to create token!",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"message": "Login succeed!",
	})
}

func Update(c *gin.Context) {
	userId := c.Param("userId")

	updateUserRequest := &models.UpdateUser{}
	if err := c.Bind(&updateUserRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Request failed!",
		})

		return
	}

	updatedUser := &models.User{}
	db := database.DB()
	db.First(&updatedUser, "id = ?", userId)
	if updatedUser.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "User not found!",
		})

		return
	}

	updatedUser.Username = updateUserRequest.Username
	updatedUser.Email = updateUserRequest.Email

	hash, err := bcrypt.GenerateFromPassword([]byte(updateUserRequest.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Failed to hash password!",
		})

		return
	}
	updatedUser.Password = string(hash)
	db.Save(updatedUser)
	
	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "User Updated Successfully!",
	})
}

func Delete(c *gin.Context) {
	userId := c.Param("userId")

	user := &models.User{}
	photo := &models.Photo{}
	db := database.DB()
	db.First(&user, "id = ?", userId)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "User not found!",
		})

		return
	}

	db.First(&photo, "user_id = ?", userId)
	if photo.UserID != 0 {
		db.Delete(photo)
	}

	db.Delete(user)
	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "User Deleted Successfully!",
	})
}