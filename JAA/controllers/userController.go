package controllers

import (
	database "JAA/databases"
	helper "JAA/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validate.new()

func HashPassword()

func VerifyPassword()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
}
