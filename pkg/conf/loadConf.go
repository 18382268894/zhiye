/**
*FileName: pkg
*Create on 2018/11/4 上午9:00
*Create by mok
*/

package conf

import (
	"github.com/spf13/viper"
	"strings"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"time"
	"os"
	"github.com/urfave/cli"
	"flag"
)

var(
	//配置文件参数
	confFile string

	//服务器配置参数
	ServerAddr string
	Runmode string
	ReadTimeOut time.Duration
	WriteTimeOut time.Duration
	PingMaxNum int

	//分页参数
	//PageSize int

	//session参数
	SessionProvider string
	SessionMaxLife int

	//redis参数
	RedisAddr string

	//jwt
	/*TokenSecret string
	TokenMaxLife int64*/

	//日志配置参数
	LogFile string
	LogLevel string
	LogFormat string

	//数据库配置参数
	DBType string
	DBUsername string
	DBPassword string
	DBAddr string
	DBName string

	//邮箱参数
	EmailHost string
	EmailPort int
	EmailAddr string
	EmailPassword string

	//手机验证配置参数
	AccessKey string
	AccessSecret string
	AliyunUrl string
	SignName string
	TemplateCode string
)



func init(){
	flag.Parse()
	c := cli.NewApp()
	c.Flags = []cli.Flag{
		cli.StringFlag{
			Name:"config",
			Usage:"config path",
		},
	}
	c.Action = func(c *cli.Context) {
		confFile = c.String("config")
	}
	err := c.Run(os.Args)
	if err != nil{
		panic(err)
	}

	viperSet(confFile)
	loadConf()
	initLog()
	logrus.Info("conf load success")
}

//使用viper加载配置文件
func loadConf(){
	ServerAddr = viper.GetString("server.addr")
	Runmode = viper.GetString("app.runmode")
	LogFile = viper.GetString("logger.file")
	LogLevel = viper.GetString("logger.level")
	LogFormat = viper.GetString("logger.format")
	PingMaxNum = viper.GetInt("app.ping_max_num")
	SessionProvider = viper.GetString("session.provider")
	SessionMaxLife = viper.GetInt("session.maxlife")
	RedisAddr = viper.GetString("redis.addr")
	//PageSize = viper.GetInt("app.pagesize")
	ReadTimeOut = viper.GetDuration("server.read_time_out") * time.Second
	WriteTimeOut = viper.GetDuration("server.write_time_out") * time.Second
	DBAddr = viper.GetString("database.addr")
	DBType = viper.GetString("database.type")
	DBUsername = viper.GetString("database.username")
	DBPassword = viper.GetString("database.password")
	DBName = viper.GetString("database.dbname")
	EmailHost = viper.GetString("email.host")
	EmailPort = viper.GetInt("email.port")
	EmailAddr = viper.GetString("email.addr")
	EmailPassword = viper.GetString("email.password")
	AccessKey = viper.GetString("aliyun.access_key")
	AccessSecret = viper.GetString("aliyun.access_secret")
	AliyunUrl = viper.GetString("aliyun.url")
	SignName = viper.GetString("aliyun.signname")
	TemplateCode = viper.GetString("aliyun.templatecode")
	/*
	TokenSecret = viper.GetString("token.secret")
	TokenMaxLife = viper.GetInt64("token.maxlife")
	*/
}


//对viper进行设置
func viperSet(confFile string)error{
	viper.SetConfigType("yaml")
	if confFile == ""{
		viper.AddConfigPath("./conf")
		viper.SetConfigName("conf")
	}
	//通过全局变量获取viper的设置
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BLOG1")
	replacer := strings.NewReplacer(".","_")
	viper.SetEnvKeyReplacer(replacer)

	err := viper.ReadInConfig()
	if err != nil{
		err = fmt.Errorf("regist viper is failed:%s",err.Error())
		return err
	}
	wathcing()
	return nil
}

//监听配置文件的改动
func wathcing(){
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logrus.Infof("config is changed:%s",e.Name)
	})
}


