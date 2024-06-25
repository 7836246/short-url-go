package config

type Seo struct {
	Title       string `mapstructure:"title" json:"title" yaml:"title"`                   // 标题
	Keywords    string `mapstructure:"keywords" json:"keywords" yaml:"keywords"`          // 关键词
	Description string `mapstructure:"description" json:"description" yaml:"description"` // 描述
}
