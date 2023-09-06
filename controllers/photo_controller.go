package controllers

import (
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/config"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/helpers"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/models"
	"gorm.io/gorm"
)

func CreatePhoto(ctx *gin.Context) {
	userId, ok := helpers.GetUserID(ctx)

	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var photo models.Photo

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	photo.UserID = userId
	err := config.DB.Create(&photo).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
			"status":  http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	ctx.Set("user", nil)

	ctx.JSON(http.StatusCreated, gin.H{
		"data_record": photo,
		"message":     "successfully create photo",
		"status":      http.StatusCreated,
	})
}

func GetPhotos(ctx *gin.Context) {
	userQuery := ctx.Query("user")

	var photos []models.Photo

	if userQuery == "" {
		config.DB.Find(&photos)
	} else {
		userId, err := strconv.Atoi(userQuery)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "bad query request",
				"status":  http.StatusBadRequest,
			})
			return
		}

		config.DB.Where("user_id = ?", userId).Find(&photos)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    photos,
		"message": "successfully fetch photos",
		"status":  http.StatusOK,
	})
}

func GetPhotoById(ctx *gin.Context) {
	var photo models.Photo

	id, err := helpers.GetParamId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad param request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if err := config.DB.First(&photo, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "photo not found",
				"status":  http.StatusNotFound,
			})
		default:
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"message": "internal server error",
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":    photo,
		"message": "successfully fetch photo",
		"status":  http.StatusOK,
	})
}

func UpdatePhoto(ctx *gin.Context) {
	userId, ok := helpers.GetUserID(ctx)

	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var photo models.Photo

	id, err := helpers.GetParamId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad param request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if _, err := govalidator.ValidateStruct(photo); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	if config.DB.Model(&photo).Where("id = ? AND user_id = ?", id, userId).Updates(&photo).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to update data",
			"status":  http.StatusBadRequest,
		})
		return
	}

	ctx.Set("user", nil)

	ctx.JSON(http.StatusOK, gin.H{
		"data_record": photo,
		"message":     "successfully update data",
		"status":      http.StatusOK,
	})
}

func DeletePhoto(ctx *gin.Context) {
	userId, ok := helpers.GetUserID(ctx)

	if !ok {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id, err := helpers.GetParamId(ctx)

	var photo models.Photo

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "bad param request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if config.DB.Unscoped().Where("id = ? AND user_id = ?", id, userId).Delete(&photo).RowsAffected == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "failed to delete data",
			"status":  http.StatusBadRequest,
		})
		return
	}

	ctx.Set("user", nil)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully delete data",
		"status":  http.StatusOK,
	})
}
