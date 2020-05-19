package config

import (
	"ddns_pro/consts"
	"fmt"
	"github.com/BurntSushi/toml"
)

var CFG *AppConfig

type Cfg struct {
	AppCfg *AppConfig `toml:"config"`
}

type AppConfig struct {
	SecretId   string   `toml:"secretId"`
	SecretKey  string   `toml:"secretKey"`
	Domain     string   `toml:"domain"`
	SubDomains []string `toml:"subDomains"`
	CheckTime  int      `toml:"checkTime"`
}

func init() {
	var configSt *Cfg
	if _, err := toml.DecodeFile(consts.ConfigFilePath, &configSt); err != nil {
		fmt.Println(err)
	}
	CFG = configSt.AppCfg
}