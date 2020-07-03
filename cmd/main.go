package main

import (
	"flag"
	"github.com/fanjq99/dnslog"
	"github.com/fanjq99/dnslog/config"
)

var (
	ymlFile = flag.String("yml", "./fixture/dev.yml", "yml config file path")
)

func main() {
	flag.Parse()
	setting, err := config.Parse(*ymlFile)
	if err != nil {
		return
	}

	dnslog.Run(setting,false)
}
