package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server/pkg/logger"
	"go-server/pkg/models"
	"go-server/pkg/redis"
	"log"
	"net/http"
	"time"
)


func GetUsers(c *gin.Context) {
	db := models.DB
	var users []models.User

	db.Find(&users)

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, users)
}

func GetUsersById(c *gin.Context) {
	db := models.DB

	id := c.Params.ByName("id")
	var user models.User

	db.First(&user, id)
	cache := redis.GetRedisCache()
	data,err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	cacheErr := cache.Set(c, id, data, 10*time.Second).Err()
	if cacheErr != nil {
		log.Println(cacheErr)
	}

	c.Header("Content-Type", "application/json")
	c.Header("Cache", "None")
	c.JSON(http.StatusOK, user)

	logger.Logger.Info("INFO: user fetched successfully")
}