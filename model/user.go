/**
*FileName: model
*Create on 2018/11/10 下午12:02
*Create by mok
*/

package model

import "time"

type User struct {
	ID uint `db:"id",json:"id"`
	Username string `db:"username",json:"username"`
	Password string `db:"password",json:"password"`
	Phone int64 `db:"phone",json:"phone"`
	Email string `db:"email",json:"email"`
	RealName string `db:"real_name",json:"real_name"`
	CreateTime time.Time `db:"create_time",json:"create_time"`
	LastIP string `db:"last_ip",json:"last_ip"`
	LastTime time.Time `db:"last_time",json:"last_time"`

}
