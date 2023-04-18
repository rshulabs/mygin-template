package core

/*
思路
1.先写好配置和全局viper
2.定义初始化方法，传入配置文件名参数（可选）
3.初始化 有一个先后顺序 传参>命令行>环境变量>默认
4.创建viper实例，设置配置文件属性（文件名）
5.读取配置，监听配置，绑定配置对象，绑定到全局viper
*/
import (
	"backup/core/internal"
	"backup/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// InitializeViper 优先级 传参>命令行>环境变量>默认
func InitializeViper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		// flag命令行
		// p *string, name string, value string, usage string) -c=value help
		flag.StringVar(&config, "c", "", "choose config file")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Println(config)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Println(config)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Println(config)
				}
			} else {
				config = configEnv
				fmt.Println(config)
			}
		} else {
			fmt.Println(config)
		}
	} else {
		config = path[0]
		fmt.Println(config)
	}
	// 创建viper实例 绑定配置
	vip := viper.New()
	vip.SetConfigFile(config)
	vip.SetConfigType("yaml")
	// 读配置
	err := vip.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file %s \n", err))
	}
	// 监控
	vip.WatchConfig()
	vip.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config changed ", in.Name)
		if err = vip.Unmarshal(&global.App.Config); err != nil {
			fmt.Println(err)
		}
	})
	// 建立映射
	if err = vip.Unmarshal(&global.App.Config); err != nil {
		fmt.Println(err)
	}
	return vip
}
