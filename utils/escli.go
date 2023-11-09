package utils

import (
	"fmt"
	
	es7 "github.com/elastic/go-elasticsearch/v7"
)

var cli *es7.Client

func MustInitESClient(conf *ESConfig) {
	var err error
	cli, err = es7.NewClient(es7.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", conf.Address, conf.Port)},
		Username:  conf.Username,
		Password:  conf.Password,
	})

	if err != nil {
		panic(err)
	}
}

func GetESCli() *es7.Client {
	return cli
}
