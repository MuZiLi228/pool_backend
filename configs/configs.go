package configs

import (
	"fmt"
	"net"
	"os"
	"pool_backend/src/util/env"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

var config = new(Config)

//Config 数据库配置
type Config struct {
	Sms struct {
		RegionID        string `yaml:"regionId"`
		AccessKeyID     string `yaml:"accessKeyId"`
		AccessKeySecret string `yaml:"accessKeySecret"`
		SignName        string `yaml:"signName"`
		TemplateCode    string `yaml:"templateCode"`
	} `yaml:"sms"`

	Pg struct {
		Read struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
			Name string `yaml:"name"`
		} `yaml:"read"`
		Write struct {
			Host string `yaml:"host"`
			Port string `yaml:"port"`
			User string `yaml:"user"`
			Pass string `yaml:"pass"`
			Name string `yaml:"name"`
		} `yaml:"write"`
		Base struct {
			MaxOpenConn     int           `yaml:"maxOpenConn"`
			MaxIdleConn     int           `yaml:"maxIdleConn"`
			ConnMaxLifeTime time.Duration `yaml:"connMaxLifeTime"`
		} `yaml:"base"`
	} `yaml:"pg"`

	//redis缓存
	Redis struct {
		Addr         string `yaml:"addr"`
		Pass         string `yaml:"pass"`
		Db           int    `yaml:"db"`
		MaxRetries   int    `yaml:"maxRetries"`
		PoolSize     int    `yaml:"poolSize"`
		MinIdleConns int    `yaml:"minIdleConns"`
	} `yaml:"redis"`

	JWT struct {
		Secret         string        `yaml:"secret"`
		ExpireDuration time.Duration `yaml:"expireDuration"`
	} `yaml:"jwt"`

	SecretKey struct {
		Aeskey string `yaml:"aeskey"`
		Iv     string `yaml:"iv"`
	} `yaml:"secretKey"`
}

//初始化consul配置文件
func init() {
	//指定读取json文件
	viper.SetConfigName(env.Active().Value() + "_configs")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("读取 配置文件报错:", err)
	}
	if err := viper.Unmarshal(config); err != nil {
		fmt.Println("解析 配置文件报错:", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			fmt.Println("监听 配置文件报错:", err)
		}
	})
}

//Get 获取配置信息
func Get() Config {
	return *config
}

//MachineID 机器id
func MachineID() string {
	return "raspi4b8g-demo-0001"
}

//ProjectName 获取项目名字
func ProjectName() string {
	return "pool_backend"
}

//ProjectHost 获取主机ip
func ProjectHost() string {
	return GetLocalIP()[0]
}

//ProjectPort 获取端口
func ProjectPort() string {
	return "7001"
}

//ConsulAddr 获取consul地址
func ConsulAddr() string {
	return fmt.Sprintf("http://%s:%s", os.Getenv("CONSUL_HOST"), os.Getenv("CONSUL_PORT"))
}

//SwaggerURL 接口文档
func SwaggerURL() string {
	return fmt.Sprintf("http://%s:%s%s", ProjectHost(), ProjectPort(), "/sys/swagger/doc.json")
}

//ProjectLogFile 日志文件目录
func ProjectLogFile() string {
	return fmt.Sprintf("./logs/%s-access.log", ProjectName())
}

//GetLocalIP 获取本地ip
func GetLocalIP() []string {
	var ipStrArr []string
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces error:", err.Error())
		return ipStrArr
	}
	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()
			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//获取IPv4
					if ipnet.IP.To4() != nil {
						ipStrArr = append(ipStrArr, ipnet.IP.String())
					}
				}
			}
		}
	}
	return ipStrArr

}
