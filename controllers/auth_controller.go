package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/config"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/helpers"
	"github.com/taufiqoo/task-5-vix-btpns-taufiqoo/models"
)

func RegisterUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if _, err := govalidator.ValidateStruct(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
			"status":  http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	var userDb models.User
	count := int64(0)

	config.DB.
		Model(&userDb).
		First(&userDb, "username = ? OR email = ?", user.Username, user.Email).
		Count(&count)

	if count > 0 {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{
			"message": "username or email is already used",
			"status":  http.StatusConflict,
		})
		return
	}

	user.Password = hashedPassword

	err = config.DB.Create(&user).Error

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "internal server error",
			"status":  http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "successfully register user",
		"status":  http.StatusCreated,
	})
}

func LoginUser(ctx *gin.Context) {
	var auth models.UserAuth

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
		})
		return
	}

	if _, err := govalidator.ValidateStruct(auth); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "bad body request",
			"status":  http.StatusBadRequest,
			"error":   err.Error(),
		})
		return
	}

	var user models.User

	if err := config.DB.First(&user, "username = ?", auth.Username).Error; err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "invalid username or password",
			"status":  http.StatusForbidden,
		})
		return
	}

	if err := helpers.CompPassword(&user.Password, &auth.Password); err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"message": "invalid username or password",
			"status":  http.StatusForbidden,
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create token",
			"status":  http.StatusInternalServerError,
			"error":   err.Error(),
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenStr, 3600*24, "", "", true, true)

	ctx.JSON(http.StatusAccepted, gin.H{
		"message": "user successfully login",
		"status":  http.StatusAccepted,
	})
}

func LogoutUser(ctx *gin.Context) {
	ctx.Set("user", nil)
	ctx.SetCookie("Authorization", "", -1, "", "", true, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully logout",
		"status":  http.StatusOK,
	})
}
