package parser

import "time"

type Parser struct {
	Id           int
	Name         string
	Description  string
	CreatedAt    time.Time
	StartParsing time.Time
	StartReport  time.Time
}
