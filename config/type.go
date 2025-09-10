package config

type Config struct {
	App appConfig `yaml:"app" validate:"required"`
}

type appConfig struct {
	Port string `yaml:"port" validate:"required"`
}
