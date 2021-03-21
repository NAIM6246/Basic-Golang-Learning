package config

type AppConfig struct {
	Secret string
}

func NewAppConfig() string {
	return "SecRet"
}
