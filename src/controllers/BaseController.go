package controllers

import "github.com/gin-gonic/gin"

type BaseController interface {
	initController(route * gin.Engine) float64
}