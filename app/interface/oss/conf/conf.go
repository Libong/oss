package conf

import (
	"flag"
	"github.com/spf13/viper"
	"libong/common/env"
	"libong/common/orm/mysql"
	"libong/common/redis"
	"libong/common/server/grpc"
	"libong/common/server/http"
	"libong/oss/app/service/oss/service/ossClient"
)

type Config struct {
	Server  *Server
	Service *Service
}

func New() *Config {
	conf := &Config{}
	//初始化配置文件
	GrabConfigFile(conf)
	return conf
}

type Service struct {
	OssClientConfig *ossClient.Config
	Dao             *Dao
}
type Dao struct {
	Mysql *commonMysql.Config
	Redis *commonRedis.Config
}
type Server struct {
	HTTP *http.Config
	GRPC *grpc.Config
}

// GrabConfigFile 获取配置文件信息
func GrabConfigFile(c *Config) {
	//获取go build 配置的程序参数用-xxx=xxx
	//flag.parse 将命令行参数进行赋值到对应的字段上 env.ConfigPath (命令行参数通过flag.set或者程序实参上设值)
	if !flag.Parsed() {
		flag.Parse()
	}
	//
	viper := viper.New()
	//设置默认的配置文件名称和后缀
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//设置配置文件的绝对路径
	viper.AddConfigPath(env.ConfigPath)
	//viper获取配置文件的最终目录是绝对路径 即/Users/qinhaokai/Documents/Projects/goForPractice/app/interface/rpcServer/cmd
	//获取工作目录 即go build 的工作目录/Users/qinhaokai/Documents/Projects/goForPractice 会拼上传入的configPath
	if err := viper.ReadInConfig(); err != nil {
		panic(err.Error())
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		panic(err)
	}
}
