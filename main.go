/**
*FileName: zhiye
*Create on 2018/11/7 下午9:57
*Create by mok
*/

package main

import (
	"zhiye/pkg/session"
	_"zhiye/pkg/conf"
)

func main(){
	mgr := session.Mgr
	ss,err := mgr.CreateSession()
	if err != nil{
		panic(err)
	}
	ss.Set("name","mok")
	ss.Set("age","24")
	err = ss.Save()
	if err != nil{
		panic(err)
	}
	defer ss.Close()
}
