package config

type App struct {
	Name   string `mapstructure:"name" json:"name" yaml:"name"`
	Env    string `mapstructure:"env" json:"env" yaml:"env"`
	Port   string `mapstructure:"port" json:"port" yaml:"port"`
	Url    string `mapstructure:"url" json:"url" yaml:"url"`
	DbType string `mapstructure:"db_type" json:"db_type" yaml:"db_type"`
}
