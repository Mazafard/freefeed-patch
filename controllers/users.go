package controllers

import (
	"feed/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type UpdateBirthdayInput struct {
	Birthday int64 `json:"birthday"`
}

// FindUser GET /books/:id
// Find a book
func FindUser(ctx *gin.Context) {
	var user models.User
	username := ctx.Param("string")

	err := models.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Content not found",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": user})

}

// UpdateUser Update a book
func UpdateBirthday(ctx *gin.Context) {
	var user models.User
	username := ctx.Param("username")
	current_user, ok := ctx.Get("UserID")
	if ok != true {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "current user not found",
		})
	}
	err := models.DB.Where("username = ? and id= ?", username, current_user).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Content not found",
		})
	}
	// Validate input
	var input UpdateBirthdayInput

	//unixTime, err := time.Parse("2006-01-02T15:04:05Z07:00")
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	timestamp := time.Unix(0, int64(input.Birthday)*int64(time.Millisecond))

	models.DB.Model(&user).Where("ID =?", current_user).Update("birthday", timestamp)
	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})

}
