package handlers

import (
	"github.com/dmonay/okra/common"
	// "fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func (dw *DoWorkResource) Register(c *gin.Context) {

	var user common.UserJson

	c.Bind(&user)

	uname := user.Username

	c.JSON(201, "You have registered, "+uname)

	err := dw.mongo.C("Users").Insert(&User{uname})
	CheckErr(err, "User not added to Users collection")
}

type DoWorkResource struct {
	mongo *mgo.Database
}

type User struct {
	Username string
}
