package config

type Server struct {
	Http Http `mapstructure:"http" json:"http" yaml:"http"` // 基础配置
	Zap  Zap  `mapstructure:"zap" json:"zap" yaml:"zap"`    // 日志配置
	Seo  Seo  `mapstructure:"seo" json:"seo" yaml:"seo"`    // 搜索引擎标题配置
}
