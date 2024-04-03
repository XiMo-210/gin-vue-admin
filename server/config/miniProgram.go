package config

type MiniProgram struct {
	Appid     string `mapstructure:"appid"`
	Secret    string `mapstructure:"secret"`
	HttpDebug bool   `mapstructure:"http_debug"`

	LogLevel     string `mapstructure:"log_level"`
	LogInfoFile  string `mapstructure:"log_info_file"`
	LogErrorFile string `mapstructure:"log_error_file"`

	State string `mapstructure:"state"`
}
