package controllers

import "github.com/gin-gonic/gin"
import (
	//"models"
	//"github.com/insionng/macross/libraries/gommon/log"
)


func Index(c *gin.Context){
	/*
	user := models.User{ UserName:"liu", LoginName:"liu", Password:"123456"}

	err := models.AddUser( &user )
	if err != nil {
		print( err.Error() )
		c.JSON(500, err)
	}

	c.JSON(200, user)
	*/

}

func  GetUsers(c *gin.Context)  {
	/*
	name := c.Query("name")
	page := models.Page{PageNum:1,PageSize:10}
	users, err:= models.GetAllUsers(name,page)
	if err == nil {
		c.JSON(200, users)
	}else{
		log.Error( err.Error() )
		c.JSON(500, err)
	}
	*/
}
