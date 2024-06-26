package parser

import (
	"context"
	"github.com/DateMine/bot-domain/pkg/models/offer_storage"
	"github.com/DateMine/bot-domain/pkg/models/request"
	"github.com/DateMine/bot-domain/pkg/models/request/clientSettings"
	"github.com/DateMine/bot-domain/pkg/models/request/requestSettings"
	"github.com/DateMine/bot-domain/pkg/models/response"
	"github.com/DateMine/bot-domain/pkg/service_provider"
	"github.com/DateMine/bot-domain/pkg/utils"
	"net/http"
	"time"
)

type ParserBase struct {
	webClient    *utils.WebClient
	Ctx          context.Context
	UseProxy     bool
	ProxyAddr    string
	Timeout      int64
	SleepTime    int64
	OfferPortion int
	OfferStorage *offer_storage.OfferStorage
}

func NewParser(serviceProvider *service_provider.ServiceProvider, ctx context.Context) *ParserBase {
	webClient := utils.NewWebClient(ctx)

	return &ParserBase{
		webClient:    webClient,
		Ctx:          ctx,
		OfferStorage: serviceProvider.OfferStorage(ctx), //TODO надо подумать
	}
}

func (p ParserBase) SendGet(url string, headers http.Header, cookies []http.Cookie, ctx context.Context) (*response.HttpClientResponse, error) {
	time.Sleep(time.Duration(p.SleepTime) * time.Microsecond)
	if headers == nil {
		headers = make(http.Header)
	}
	if cookies == nil {
		cookies = make([]http.Cookie, 0)
	}
	response, err := p.webClient.Send(request.Request{
		RequestSettings: requestSettings.RequestSettings{
			Url:     url,
			Method:  requestSettings.GET,
			Headers: headers,
			Cookies: cookies,
		},
		ClientSettings: clientSettings.ClientSettings{
			UseProxy: p.UseProxy,
			Proxy:    p.ProxyAddr,
			TimeOut:  p.Timeout,
		},
	}, p.Ctx)
	if err != nil {
		return nil, err
	}
	return response, err
}

func (p ParserBase) SendPost(url string, body string, headers http.Header, cookies []http.Cookie, ctx context.Context) (*response.HttpClientResponse, error) {
	time.Sleep(time.Duration(p.SleepTime) * time.Microsecond)
	if headers == nil {
		headers = make(http.Header)
	}
	if cookies == nil {
		cookies = make([]http.Cookie, 0)
	}
	response, err := p.webClient.Send(request.Request{
		RequestSettings: requestSettings.RequestSettings{
			Url:     url,
			Method:  requestSettings.POST,
			Headers: headers,
			Cookies: cookies,
			Body:    body,
		},
		ClientSettings: clientSettings.ClientSettings{
			UseProxy: p.UseProxy,
			Proxy:    p.ProxyAddr,
			TimeOut:  p.Timeout,
		},
	}, p.Ctx)
	if err != nil {
		return nil, err
	}

	return response, err
}
