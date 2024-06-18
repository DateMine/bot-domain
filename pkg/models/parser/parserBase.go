package parser

import (
	"github.com/DateMine/bot-domain/pkg/utils"
)

type ParserBase struct {
	webClient utils.WebClient
}

func (p ParserBase) Start() error {
	return nil
}

func (p ParserBase) GenerateReport() error {
	return nil
}
