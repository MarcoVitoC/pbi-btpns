package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/MarcoVitoC/pbi-btpns/database"
	"github.com/MarcoVitoC/pbi-btpns/models"
)

func UploadPhoto(c *gin.Context) {
	loggedInUser, _ := c.Get("user")
	user, _ := loggedInUser.(*models.User)

	photoRequest := &models.Photo{}
	if err := c.Bind(&photoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Request failed!",
		})

		return
	}

	newPhoto := models.Photo{
		Title: photoRequest.Title, 
		Caption: photoRequest.Caption, 
		PhotoUrl: photoRequest.PhotoUrl, 
		UserID: user.ID,
	}

	result := database.DB().Create(&newPhoto)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": result.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "Photo Uploaded Successfully!",
	})
}

func GetPhoto(c *gin.Context) {
	var photos []models.Photo

	result := database.DB().Find(&photos)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Failed to fetch all photos!",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "Photos Fetched Successfully!",
		"result": photos,
	})
}

func UpdatePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	updatePhotoRequest := &models.UpdatePhoto{}
	if err := c.Bind(&updatePhotoRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Request failed!",
		})

		return
	}

	updatedPhoto := &models.Photo{}
	db := database.DB()
	db.First(&updatedPhoto, "id = ?", photoId)
	if updatedPhoto.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Photo not found!",
		})

		return
	}

	updatedPhoto.Title = updatePhotoRequest.Title
	updatedPhoto.Caption = updatePhotoRequest.Caption
	updatedPhoto.PhotoUrl = updatePhotoRequest.PhotoUrl
	db.Save(updatedPhoto)
	
	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "Photo Updated Successfully!",
	})
}

func DeletePhoto(c *gin.Context) {
	photoId := c.Param("photoId")

	photo := &models.Photo{}
	db := database.DB()
	db.First(&photo, "id = ?", photoId)
	if photo.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H {
			"code": 400,
			"message": "Photo not found!",
		})

		return
	}

	db.Delete(photo)
	c.JSON(http.StatusOK, gin.H {
		"code": 200,
		"message": "Photo Deleted Successfully!",
	})
}