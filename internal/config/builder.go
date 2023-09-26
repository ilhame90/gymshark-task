package config

import (
	"strconv"
	"strings"

	"slices"

	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

// Config is the main config holder
type Config struct {
	Env              string
	HttpServerConfig HttpServerConfig
	Log              LoggingConfig
	Packs            []int
}

type HttpServerConfig struct {
	CORSConfig middleware.CORSConfig
	Port       string
}

type LoggingConfig struct {
	Level string
}

// New is for config building
func New() *Config {
	configBuilder := viper.New()
	configBuilder.SetEnvPrefix("orders")
	configBuilder.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	configBuilder.SetEnvKeyReplacer(replacer)

	config := &Config{
		Env: configBuilder.GetString("env"),
		Log: LoggingConfig{
			Level: configBuilder.GetString("log.level"),
		},
		HttpServerConfig: HttpServerConfig{
			CORSConfig: middleware.CORSConfig{
				AllowOrigins: strings.Split(configBuilder.GetString("http.alloworigins"), ","),
				AllowHeaders: strings.Split(configBuilder.GetString("http.allowheaders"), ","),
			},
			Port: configBuilder.GetString("port"),
		},
		Packs: normalizePacks(configBuilder.GetString("packs")),
	}

	return config
}

func normalizePacks(packsStr string) []int {
	packList := strings.Split(packsStr, ",")
	packs := make([]int, len(packList))
	for i, packStr := range packList {
		p, err := strconv.Atoi(packStr)
		if err != nil {
			panic(err)
		}

		packs[i] = p
	}

	slices.Sort(packs)
	return packs
}
