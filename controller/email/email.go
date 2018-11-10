/**
*FileName: email
*Create on 2018/11/10 下午7:25
*Create by mok
*/

package email

import (
	"github.com/gin-gonic/gin"
	"zhiye/pkg/token"
	"net/http"
	"zhiye/model"
)

//注册时激活账号
func CheckEmail(c *gin.Context){
	tokenss := c.Query("token")
	claim,err := token.ParseToken(tokenss)
	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Code":400,
			"Message":err.Error(),
		})
		return
	}
	var user model.User
	user.Username =  claim.Username
	err = user.Active()
	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"Code":400,
			"Message":err.Error(),
		})
		return
	}
	c.JSON(http.StatusBadRequest,gin.H{
		"Code":200,
		"Message":"OK",
	})
}