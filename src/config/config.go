package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	App   App   `toml:"app"`
	Db    Db    `toml:"db"`
	Redis Redis `toml:"redis"`
}

type App struct {
	Port                        string `toml:"port"`
	TimeoutToGracefulShutdownMs int    `toml:"timeoutToGracefulShutdownMs"`
	CorsOrigins                 string `toml:"corsOrigins"`
}

type Db struct {
	User                      string `toml:"user"`
	Name                      string `toml:"name"`
	Password                  string `toml:"password"`
	Host                      string `toml:"host"`
	Port                      string `toml:"port"`
	MaxOpenConnection         int    `toml:"maxOpenConnection"`
	MaxIdleConnection         int    `toml:"maxIdleConnection"`
	ConnectionMaxLifetimeHour int    `toml:"connectionMaxLifetimeHour"`
}

type Redis struct {
	URL                string `toml:"url"`
	Password           string `toml:"password"`
	DB                 int    `toml:"db"`
	PoolSizePerCPU     int    `toml:"poolSizePerCPU"`
	MinIdleConnection  int    `toml:"minIdleConnection"`
	PoolTimeoutSeconds int    `toml:"poolTimeoutSeconds"`
}

var Conf *Config

const repositoryName = "go-clean-architecture"

func getProjectRootPath() string {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, repositoryName) && !strings.HasSuffix(wd, "app") {
		wd = filepath.Dir(wd)
	}

	return wd
}

func init() {
	Conf = new(Config)

	GO_ENV := os.Getenv("GO_ENV")
	if GO_ENV == "" {
		GO_ENV = "development"
	}

	viper.SetConfigName(GO_ENV)
	viper.AddConfigPath(getProjectRootPath() + "/src/config")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %s \n", err))
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("failed to unmarshal err: %s \n", err))
	}
}

func IsDev() bool {
	return os.Getenv("GO_ENV") == "development"
}

func IsStg() bool {
	return os.Getenv("GO_ENV") == "staging"
}

func IsProd() bool {
	return os.Getenv("GO_ENV") == "production"
}

func IsTest() bool {
	return os.Getenv("GO_ENV") == "test"
}
