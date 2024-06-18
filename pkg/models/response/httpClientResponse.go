package response

import (
	"golang.org/x/net/html"
)

type HttpClientResponse struct {
	Content    string
	ByteData   []byte
	StatusCode int64
	Html       *html.Node
	//Todo добавить куки и хедеры
}
