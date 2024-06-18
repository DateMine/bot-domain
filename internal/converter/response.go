package converter

import (
	"bytes"
	respDomain "github.com/DateMine/bot-domain/pkg/models/response"
	"github.com/DateMine/grpc-domain/pkg/parser_v1"
	"golang.org/x/net/html"
)

func ToResponseDomain(response *parser_v1.ParsingResponse) *respDomain.HttpClientResponse {
	body := response.Body
	content := string(body)
	r := bytes.NewReader(body)

	htmlDoc, err := html.Parse(r)
	if err != nil {
		htmlDoc = nil
	}
	res := &respDomain.HttpClientResponse{
		Content:    content,
		ByteData:   body,
		StatusCode: response.StatusCode,
		Html:       htmlDoc,
	}
	return res
}
