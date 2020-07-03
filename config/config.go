package config

import (
	"io/ioutil"

	"github.com/fanjq99/common/log"
	"gopkg.in/yaml.v2"
)

type YmlConfig struct {
	DnsDomain string      `yaml:"dns_domain"`
	ServerIp  string      `yaml:"server_ip"`
	SaveTime  int64       `yaml:"save_time"`
	ApiAddr   string      `yaml:"api_addr"`
	Redis     RedisConfig `yaml:"redis,flow"`
}

type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

func Parse(file string) (YmlConfig, error) {
	var yc YmlConfig
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Error(err)
		return yc, err
	}

	err = yaml.Unmarshal(body, &yc)
	if err != nil {
		log.Error(err)
		return yc, err
	}

	return yc, nil
}
