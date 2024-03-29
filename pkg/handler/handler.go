package handler

import "github.com/gin-gonic/gin"

func CreateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "CreateUser"})
}

func DeleteUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "DeleteUser"})
}

func UpdateUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "UpdateUser"})
}

func GetUser(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetUser"})
}

func GetUsers(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "GetUsers"})
}
