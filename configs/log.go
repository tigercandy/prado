package configs

type Log struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Path          string `mapstructure:"path" json:"path" yaml:"path"`
	Filename      string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine      bool `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups    int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize       int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxAge        int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress      bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
	ConsoleStdout bool   `mapstructure:"console_stdout" json:"console_stdout" yaml:"console_stdout"`
	FileStdout    bool   `mapstructure:"file_stdout" json:"file_stdout" yaml:"file_stdout"`
	LocalTime     bool   `mapstructure:"local_time" json:"local_time" yaml:"local_time"`
}
