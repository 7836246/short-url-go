package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
	"short-url-go/internal/global"
	"short-url-go/internal/initialize/internal"
)

// Viper 函数用于初始化 Viper 实例，并加载配置文件。
// 配置加载优先级: 命令行参数 > 环境变量 > 默认值
// 作者 [SliverHorn](https://github.com/SliverHorn)
func Viper(path ...string) *viper.Viper {
	var config string

	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "选择配置文件路径.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" { // 判断 internal.ConfigEnv 常量存储的环境变量是否为空
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDefaultFile
					fmt.Printf("当前 Gin 模式为%s，使用默认配置文件路径：%s\n", gin.Mode(), internal.ConfigDefaultFile)
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
					fmt.Printf("当前 Gin 模式为%s，使用发布配置文件路径：%s\n", gin.Mode(), internal.ConfigReleaseFile)
				case gin.TestMode:
					config = internal.ConfigTestFile
					fmt.Printf("当前 Gin 模式为%s，使用测试配置文件路径：%s\n", gin.Mode(), internal.ConfigTestFile)
				}
			} else { // internal.ConfigEnv 常量存储的环境变量不为空，使用环境变量中指定的配置文件路径
				config = configEnv
				fmt.Printf("使用环境变量%s指定的配置文件路径：%s\n", internal.ConfigEnv, config)
			}
		} else { // 使用命令行参数中指定的配置文件路径
			fmt.Printf("使用命令行参数 -c 指定的配置文件路径：%s\n", config)
		}
	} else { // 使用函数传递的第一个参数作为配置文件路径
		config = path[0]
		fmt.Printf("使用 Viper 函数传递的配置文件路径：%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("读取配置文件发生致命错误：%s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件发生变更:", e.Name)
		if err = v.Unmarshal(&global.G_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.G_CONFIG); err != nil {
		panic(err)
	}

	return v
}
