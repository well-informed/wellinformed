package wellinformed

import (
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	ServerPort string    `split_words:"true"`
	DBHost     string    `split_words:"true"`
	DBName     string    `split_words:"true"`
	DBUser     string    `split_words:"true"`
	DBPassword string    `split_words:"true"`
	LogLevel   log.Level `split_words:"true"`
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
	} else if env.Env == "unittest" {
		conf = unitTestConfig
	} else {
		conf = devConfig
	}
	envconfig.MustProcess("wellinformed", &conf)
	return conf

}

var devConfig = Config{
	ServerPort: "8080",
	DBHost:     "localhost",
	DBName:     "postgres",
	DBUser:     "postgres",
	DBPassword: "password",
	LogLevel:   log.DebugLevel,
}

var unitTestConfig = Config{
	ServerPort: "8081",
	DBHost:     "localhost",
	DBName:     "unittest",
	DBUser:     "postgres",
	DBPassword: "password",
	LogLevel:   log.DebugLevel,
}

var prodConfig = Config{
	ServerPort: "80",
	DBHost:     "edyn.c7xblzysdvfi.us-east-2.rds.amazonaws.com",
	DBName:     "edyn",
	DBUser:     "edyn",
	DBPassword: "MPyDqCs4NCcCRe",
	LogLevel:   log.InfoLevel,
}
