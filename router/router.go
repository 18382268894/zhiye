/**
*FileName: router
*Create on 2018/11/10 上午11:25
*Create by mok
*/

package router

import (
	"github.com/gin-gonic/gin"
	"zhiye/controller/user"

)

func LoadRouter(router *gin.Engine){
	userG := router.Group("/user")
	{
		userG.POST("",user.Regist)  //用户注册
	}
}
