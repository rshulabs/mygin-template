package config

type Jwt struct {
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"`
	Expire int64  `mapstructure:"expire" json:"expire" yaml:"expire"`
}
