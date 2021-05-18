package main

import (
	"fmt"
	"github.com/wizard7414/epos_v2/utils"
	"strconv"
)

func main() {
	configuration := utils.GetConfig("config/deviant-conf.json")

	fmt.Println(strconv.Itoa(configuration.ParseLimit))
}
