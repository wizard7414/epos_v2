package parser

import "github.com/wizard7414/epos_v2/domain/miner"

type WebParser interface {
	ParseFromUrl(pathUrl string) (miner.ObjectV, error)

	ParseLinksFromUrl(pathUrl string) ([]string, error)
}

type FileParser interface {
	ParseFromFile(pathFile string) (miner.ObjectV, error)

	ParseLinksFromFile(pathFile string) ([]string, error)
}
