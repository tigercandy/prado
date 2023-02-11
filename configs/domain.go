package configs

type Domain struct {
	GetHost bool   `mapstructure:"get_host" json:"get_host" yaml:"get_host"`
	Url     string `mapstructure:"url" json:"url" yaml:"url"`
}
