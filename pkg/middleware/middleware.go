package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-server/pkg/models"
	"go-server/pkg/redis"
	"log"
	"net/http"
)

func VerifyCache()  gin.HandlerFunc {
	return func(c *gin.Context) {
		cache := redis.GetRedisCache()
		id := c.Params.ByName("id")

		data, err := cache.Get(c, id).Bytes()
		if err != nil {
			c.Next()
			return
		}

		var user  models.User
		err = json.Unmarshal(data,&user)
		if err != nil {
			log.Println(err)
		}

		c.Header("Cache", "Hit")
		c.JSON(http.StatusOK, user)
		c.Abort()
		return
	}
}