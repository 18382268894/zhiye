/**
*FileName: model
*Create on 2018/11/10 下午12:02
*Create by mok
*/

package model

import (
	"time"
	"database/sql"
	"github.com/pkg/errors"
	"strings"
)

type User struct {
	ID uint `db:"id",json:"id"`
	Username string `db:"username",json:"username"`
	Password string `db:"password",json:"password"`
	Phone sql.NullInt64 `db:"phone",json:"phone"`
	Email sql.NullString `db:"email",json:"email"`
	RealName string `db:"real_name",json:"real_name"`
	CreateTime time.Time `db:"create_time",json:"create_time"`
	LastIP string `db:"last_ip",json:"last_ip"`  //最后一次登录IP
	LastTime time.Time `db:"last_time",json:"last_time"` //最后一次登录时间
	Status int `db:"status",json:"status"`
}


var(
	NameDuplicate = errors.New("用户名已经存在")
	EmailDuplicate = errors.New("该邮箱已经注册")
	PhoneDuplicate = errors.New("该手机号码已经被注册")
)

//创建用户
func (u *User)Create()(error){
	sqlstr := `insert into users(username,password,phone,email) values(?,?,?,?)`
	_,err := DB.Exec(sqlstr,u.Username,u.Password,u.Phone,u.Email)
	if err != nil{
		if strings.Index(err.Error(),"Duplicate entry") != -1{
			switch  {
			case strings.LastIndex(err.Error(),`'username'`) !=-1:
				return NameDuplicate
			case strings.LastIndex(err.Error(),`'phone'`) !=-1:
				return PhoneDuplicate
			case strings.LastIndex(err.Error(),`'email'`) !=-1:
				return EmailDuplicate
			}
		}
		return err
	}
 	return nil
}

//账号激活
func (u *User)Active()error{
	sqlstr := `update users set status=? where username=?`
	_,err := DB.Exec(sqlstr,1,u.Username)
	if err != nil{
		return err
	}
	return nil
}