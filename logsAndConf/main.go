package main

import (
	"fmt"
	"mysqlTabelMerger/pkg/conf"
	"mysqlTabelMerger/pkg/log"
	"time"
)

func main()  {
	cfg, err := conf.Init("config/config.yaml")
	fmt.Println(cfg)
	if err != nil {
		panic(err)
	}
	log.Init(&cfg.Logger)
	log.Info("asdasds")
	time.Sleep(1)
}