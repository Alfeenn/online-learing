package controller

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(g *gin.Context)
	Update(g *gin.Context)
	Delete(g *gin.Context)
	Find(g *gin.Context)
	FindAll(g *gin.Context)
	UserSignIn(g *gin.Context)
	GetAccessList(g *gin.Context)
}
