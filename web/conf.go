package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func InitConfig(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initConfig() error {

	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 如果指定了配置文件，则解析指定的配置文件
	} else {
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yml") // 设置配置文件格式为YML

	if err := viper.ReadInConfig(); err != nil { // viper解析配置文件
		return err
	}

	return nil
}
