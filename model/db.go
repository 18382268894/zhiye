/**
*FileName: model
*Create on 2018/11/10 下午1:08
*Create by mok
*/

package model

import (
	"github.com/jmoiron/sqlx"
	_"github.com/go-sql-driver/mysql"

	"fmt"
	"zhiye/pkg/conf"
	"time"
)

var DB *sqlx.DB
func Init(){
	var err error
	dsn :=fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		conf.DBUsername,conf.DBPassword,conf.DBAddr,conf.DBName)
	DB,err = sqlx.Open(conf.DBType,dsn)
	if err != nil{
		panic(err)
	}
	DB.SetConnMaxLifetime(100*time.Second)
	DB.SetMaxIdleConns(600)
	DB.SetMaxOpenConns(5000)
}
func Close(){
	DB.Close()
}