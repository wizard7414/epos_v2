package utils

import (
	"github.com/tkanos/gonfig"
	"github.com/wizard7414/epos_v2/domain"
)

func GetConfig(configPath string) *domain.EposConfig {
	configuration := domain.EposConfig{}
	err := gonfig.GetConf(configPath, &configuration)
	if err != nil {
		panic(err)
	}

	return &configuration
}
