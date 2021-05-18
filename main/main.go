package main

import (
	"fmt"
	"github.com/tkanos/gonfig"
	"github.com/wizard7414/epos_v2/domain"
	"strconv"
)

func main() {
	configuration := domain.EposConfig{}
	err := gonfig.GetConf("config/deviant-conf.json", &configuration)
	if err != nil {
		panic(err)
	}
	fmt.Println(strconv.Itoa(configuration.ParseLimit))
}
