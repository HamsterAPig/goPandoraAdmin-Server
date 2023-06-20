package config

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type variableConfig struct {
	Listen       string `mapstructure:"listen"`      // 监听地址
	ProxyNode    string `mapstructure:"proxy"`       // 代理地址
	DatabasePath string `mapstructure:"database"`    // 数据库路径
	DebugLevel   string `mapstructure:"debug-level"` // 日志等级
}

var Conf = new(variableConfig)

func ReadConfig() (*variableConfig, error) {
	cmdViper := viper.New()
	fileViper := viper.New()

	pflag.StringP("config", "c", "", "config file path")
	pflag.StringP("listen", "l", "", "listen address")
	pflag.StringP("proxy", "p", "", "proxy address")
	pflag.StringP("database", "d", "./goPandora.db", "database path")

	err := cmdViper.BindPFlags(pflag.CommandLine)
	if err != nil {
		return nil, fmt.Errorf("failed to bind flags: %w", err)
	}
	configFilePath := cmdViper.GetString("config")
	if configFilePath != "" {
		fileViper.SetConfigFile(configFilePath)
		fileViper.SetConfigType("yaml")

		if err := fileViper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		err := cmdViper.MergeConfigMap(fileViper.AllSettings())
		if err != nil {
			return nil, fmt.Errorf("failed to merge config: %w", err)
		}
	}
	var ret = new(variableConfig)
	if err := cmdViper.Unmarshal(ret); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	return ret, nil
}
