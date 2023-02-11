package configs

type Configuration struct {
	App      App      `mapstructure:"app" json:"app" yaml:"app"`
	Domain   Domain   `mapstructure:"domain" json:"domain" yaml:"domain"`
	Database Database `mapstructure:"db" json:"db" yaml:"db"`
	Redis    Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log      Log      `mapstructure:"log" json:"log" yaml:"log"`
	Jwt      Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	File     File     `mapstructure:"file" json:"file" yaml:"file"`
	Mail     Mail     `mapstructure:"mail" json:"mail" yaml:"mail"`
}
