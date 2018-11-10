/**
*FileName: email
*Create on 2018/11/10 下午3:06
*Create by mok
*/

package email

import (
	"github.com/go-gomail/gomail"
	"zhiye/pkg/conf"
	"zhiye/pkg/token"
	"fmt"
	"github.com/sirupsen/logrus"
)


func SendTo(username,email string)error{
	m := gomail.NewMessage()
	m.SetHeader("To",email)
	m.SetAddressHeader("From","1005914310@qq.com","之也")
	m.SetHeader("Subject","账号注册验证")
	logrus.Info("username1:",username)
	tokenss,err := token.NewToken(username)
	if err !=nil{
		return err
	}
	Context := `
<html>
	<head>
		<title>之也邮箱验证</title>
	</head>
	<body>
		<p><a href="http://127.0.0.1:8080/email?token=`+tokenss+`">点击激活账号</a></p>
	</body>
</html>
`
	m.SetBody("text/html",Context)
	d := gomail.NewDialer(conf.EmailHost,conf.EmailPort,conf.EmailAddr,conf.EmailPassword)
	err = d.DialAndSend(m)
	if err != nil{
		return fmt.Errorf("send email err:%s",err.Error())
	}
	return nil
}