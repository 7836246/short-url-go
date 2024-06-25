package config

type Http struct {
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	RouterPrefix string `mapstructure:"routerPrefix" json:"routerPrefix" yaml:"routerPrefix"`
	DbName       string `mapstructure:"dbName" json:"dbName" yaml:"dbName"`
}
