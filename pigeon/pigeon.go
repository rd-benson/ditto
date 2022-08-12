package pigeon

import "fmt"

var cfg Config

func Start(cfgFile string) {
	cfg.ReadInConfig(cfgFile)
	fmt.Println(cfg)
}
