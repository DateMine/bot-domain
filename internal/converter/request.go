package converter

import (
	"github.com/DateMine/bot-domain/pkg/models/request"
	"github.com/DateMine/bot-domain/pkg/models/request/clientSettings"
	"github.com/DateMine/bot-domain/pkg/models/request/requestSettings"
	"github.com/DateMine/grpc-domain/pkg/parser_v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToRequestDesc(request request.Request) *parser_v1.ParsingRequest {
	return &parser_v1.ParsingRequest{
		HttpClientSettings: ToHttpClientSettingsDesc(request.ClientSettings),
		RequestSettings:    ToRequestSettings(request.RequestSettings),
	}
}

func ToHttpClientSettingsDesc(settings clientSettings.ClientSettings) *parser_v1.HttpClientSettings {
	var proxy = wrapperspb.String(settings.Proxy)
	return &parser_v1.HttpClientSettings{
		Proxy:    proxy,
		UseProxy: settings.UseProxy,
		Timeout:  settings.TimeOut,
	}
}

func ToRequestSettings(settings requestSettings.RequestSettings) *parser_v1.RequestSettings {
	var body = wrapperspb.String(settings.Body)
	return &parser_v1.RequestSettings{
		Url:     settings.Url,
		Method:  ToMethodDesc(settings.Method),
		Body:    body,
		Headers: make([]*parser_v1.Header, 0),
	}
}

func ToMethodDesc(m requestSettings.Method) parser_v1.Method {
	switch m {
	case requestSettings.GET:
		return parser_v1.Method_get
	case requestSettings.POST:
		return parser_v1.Method_post
	default:
		return parser_v1.Method_get
	}
}
