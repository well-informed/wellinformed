package wellinformed

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort           string `split_words:"true"`
	DBHost               string `split_words:"true"`
	DBName               string `split_words:"true"`
	DBUser               string `split_words:"true"`
	DBPassword           string `split_words:"true"`
	DBMaxOpenConnections int
	DBMaxIdleConnections int
	LogLevel             log.Level `split_words:"true"`
}

type Env struct {
	Env string `split_words:"true" default:"dev"`
}

func GetConfig() Config {
	var env Env
	var conf Config

	envconfig.MustProcess("wellinformed", &env)
	if env.Env == "production" {
		conf = prodConfig
	} else {
		conf = devConfig
	}
	envconfig.MustProcess("wellinformed", &conf)
	return conf

}

var devConfig = Config{
	ServerPort:           "8080",
	DBHost:               "localhost",
	DBName:               "postgres",
	DBUser:               "postgres",
	DBPassword:           "password",
	DBMaxOpenConnections: 50,
	DBMaxIdleConnections: 10,
	LogLevel:             log.DebugLevel,
}

var prodConfig = Config{
	ServerPort:           "80",
	DBHost:               "edyn.c7xblzysdvfi.us-east-2.rds.amazonaws.com",
	DBName:               "edyn",
	DBUser:               "edyn",
	DBPassword:           "MPyDqCs4NCcCRe",
	DBMaxOpenConnections: 50,
	DBMaxIdleConnections: 10,
	LogLevel:             log.InfoLevel,
}
