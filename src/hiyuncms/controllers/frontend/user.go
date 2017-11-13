package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/yy"
	"fmt"
	"hiyuncms/controllers/backend"
)

func UserLoginShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"path":"",
	})
}

func UserLogin(c * gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"path":"",
	})
}

func RegistryShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"path":"",
		"regsuccess":false,
	})
}

func Registry(c * gin.Context)  {
	user := yy.YyUser{}

	company := yy.YyCompany{}
	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", company)
	c.Bind( &user )
	c.Bind( &company )

	fmt.Printf("%v\n", user)
	fmt.Printf("%v\n", company)

	user.UserPassword = backend.Md5str(user.UserPassword)

	err := yy.CompanyReg(&company, &user)
	if err != nil {
		fmt.Printf("%s\n", err.Error() )
	}
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"path":"",
		"regsuccess":true,
	})
}