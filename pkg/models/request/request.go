package request

import (
	cs "github.com/DateMine/bot-domain/pkg/models/request/clientSettings"
	rs "github.com/DateMine/bot-domain/pkg/models/request/requestSettings"
)

type Request struct {
	RequestSettings rs.RequestSettings
	ClientSettings  cs.ClientSettings
}
